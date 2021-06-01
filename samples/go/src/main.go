package main

import (
	"fmt"
	"net/http"

	"leadIntegration/src/controllers"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/health-check", controllers.HealthCheck).Methods("GET")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err)
	}
}
