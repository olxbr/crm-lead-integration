package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"leadIntegration/src/models"
	"leadIntegration/src/utils"
)

var HealthCheck = func(w http.ResponseWriter, r *http.Request) {
	utils.Message(w, "OK")
}

var RecieveLead = func(w http.ResponseWriter, r *http.Request) {

	lead := models.Lead{}
	err := json.NewDecoder(r.Body).Decode(&lead)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		utils.Message(w, "Request Invalido")
		return
	}

	utils.Message(w, "Lead successfully received")
}
