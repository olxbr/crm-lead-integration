package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"leadIntegration/src/auth"
	"leadIntegration/src/controllers"
)

func TestHealthCheck(t *testing.T) {

	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.HealthCheck)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"message": "OK"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestRecieveLead(t *testing.T) {

	requestBody, err := json.Marshal(map[string]string{
		"leadOrigin":      "VivaReal",
		"timestamp":       "2017-10-23T15:50:30.619Z",
		"originLeadId":    "59ee0fc6e4b043e1b2a6d863",
		"originListingId": "87027856",
		"clientListingId": "a40171",
		"name":            "Nome Consumidor",
		"email":           "nome.consumidor@email.com",
		"ddd":             "11",
		"phone":           "999999999",
		"phoneNumber":     "11999999999",
		"message":         "Olá, tenho interesse neste imóvel. Aguardo o contato. Obrigado."})

	req, err := http.NewRequest("POST", "/leads/lead", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.SetBasicAuth("vivareal", os.Getenv("SECRET_KEY"))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.RecieveLead)

	handler.ServeHTTP(rr, req)

	user, secretKey, _ := req.BasicAuth()

	if !auth.ValidAuthorization(user, secretKey) {
		t.Errorf("handler returned wrong authorization")
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"message": "Lead successfully received"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestWrongAuthorization(t *testing.T) {

	requestBody, err := json.Marshal(map[string]string{
		"leadOrigin":      "VivaReal",
		"timestamp":       "2017-10-23T15:50:30.619Z",
		"originLeadId":    "59ee0fc6e4b043e1b2a6d863",
		"originListingId": "87027856",
		"clientListingId": "a40171",
		"name":            "Nome Consumidor",
		"email":           "nome.consumidor@email.com",
		"ddd":             "11",
		"phone":           "999999999",
		"phoneNumber":     "11999999999",
		"message":         "Olá, tenho interesse neste imóvel. Aguardo o contato. Obrigado."})

	req, err := http.NewRequest("POST", "/leads/lead", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.SetBasicAuth("user", "password")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.RecieveLead)

	handler.ServeHTTP(rr, req)

	user, secretKey, _ := req.BasicAuth()

	if auth.ValidAuthorization(user, secretKey) {
		t.Errorf("handler returned wrong authorization")
	}

}
