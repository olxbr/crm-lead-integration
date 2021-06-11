package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"leadIntegration/src/controllers"

	"github.com/joho/godotenv"
)

var bodyTest = []struct {
	test string
	body map[string]string
}{
	{
		test: "Vivareal",
		body: map[string]string{
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
			"message":         "Olá, tenho interesse neste imóvel. Aguardo o contato. Obrigado.",
		},
	},
	{
		test: "Zap",
		body: map[string]string{
			"leadOrigin":      "Zap",
			"timestamp":       "2017-10-23T15:50:30.619Z",
			"originLeadId":    "59ee0fc6e4b043e1b2a6d863",
			"originListingId": "87027856",
			"clientListingId": "a40171",
			"name":            "Nome Consumidor",
			"email":           "nome.consumidor@email.com",
			"ddd":             "21",
			"phone":           "888888888",
			"message":         "Olá, tenho interesse neste imóvel. Aguardo o contato. Obrigado.",
		},
	},
	{
		test: "OLX",
		body: map[string]string{
			"leadOrigin":      "OLX",
			"timestamp":       "2017-10-23T15:50:30.619Z",
			"originLeadId":    "59ee0fc6e4b043e1b2a6d863",
			"originListingId": "87027856",
			"clientListingId": "a40171",
			"name":            "Nome Consumidor",
			"email":           "nome.consumidor@email.com",
			"ddd":             "31",
			"phone":           "777777777",
			"message":         "Olá, tenho interesse neste imóvel. Aguardo o contato. Obrigado.",
		},
	},
}

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

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	for _, tt := range bodyTest {
		t.Run(tt.test, func(t *testing.T) {
			requestBody, err := json.Marshal(tt.body)

			req, err := http.NewRequest("POST", "/leads/lead", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}

			req.SetBasicAuth("vivareal", os.Getenv("SECRET_KEY"))

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(controllers.RecieveLead)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			expected := `{"message": "Lead successfully received"}`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		})
	}

}

func TestWrongAuthorization(t *testing.T) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	for _, tt := range bodyTest {
		t.Run(tt.test, func(t *testing.T) {
			requestBody, err := json.Marshal(tt.body)

			req, err := http.NewRequest("POST", "/leads/lead", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}

			req.SetBasicAuth("user", "password")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(controllers.RecieveLead)

			handler.ServeHTTP(rr, req)

			expected := `{"message": "Authorization header invalid"}`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}

		})
	}
}
