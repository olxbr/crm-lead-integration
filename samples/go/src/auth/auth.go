package auth

import (
	"net/http"
	"os"

	"leadIntegration/src/utils"
)

func ValidAuthorization(user string, secretKey string) bool {
	return user == "vivareal" && secretKey == os.Getenv("SECRET_KEY")
}

var BasicAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		noAuth := []string{"/health-check"}
		requestPath := r.URL.Path

		for _, value := range noAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		user, secretKey, _ := r.BasicAuth()

		if !ValidAuthorization(user, secretKey) {

			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, "Authorization header invalid")
			return
		}

		next.ServeHTTP(w, r)
	})
}
