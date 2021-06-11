package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"leadIntegration/src/models"
	"leadIntegration/src/utils"
)

var HealthCheck = func(w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, "OK")
}

var RecieveLead = func(w http.ResponseWriter, r *http.Request) {

	user, secretKey, _ := r.BasicAuth()

	if user != "vivareal" || secretKey != os.Getenv("SECRET_KEY") {
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, "Authorization header invalid")
		return
	}

	lead := models.Lead{}
	err := json.NewDecoder(r.Body).Decode(&lead)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, "Invalid request")
		return
	}

	utils.Respond(w, "Lead successfully received")
}
