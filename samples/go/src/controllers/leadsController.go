package controllers

import (
	"net/http"

	utils "leadIntegration/src/utils"
)

var HealthCheck = func(w http.ResponseWriter, r *http.Request) {
	utils.Message(w, "OK")
}
