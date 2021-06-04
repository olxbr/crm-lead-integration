package utils

import (
	"fmt"
	"io"
	"net/http"
)

func Respond(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, fmt.Sprintf(`{"message": "%s"}`, message))

}
