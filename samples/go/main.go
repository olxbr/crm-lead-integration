package main

import (
	"log"
	"net/http"

	"leadIntegration/src/auth"
	"leadIntegration/src/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(err.Error())
	}

	router := mux.NewRouter()

	router.Use(auth.BasicAuthentication)

	router.HandleFunc("/health-check", controllers.HealthCheck).Methods("GET")
	router.HandleFunc("/leads/lead", controllers.RecieveLead).Methods("POST")

	log.Print("Server up on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
