package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"leadIntegration/src/models"
	"leadIntegration/src/utils"
)

var HealthCheck = func(w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, "OK")
}

var RecieveLead = func(w http.ResponseWriter, r *http.Request) {

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
