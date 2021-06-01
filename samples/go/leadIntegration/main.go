package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeLink)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err)
	}
}
