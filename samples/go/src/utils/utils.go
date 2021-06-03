package utils

import (
	"fmt"
	"io"
	"net/http"
)

func Message(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, fmt.Sprintf(`{"message": "%s"}`, message))

}
