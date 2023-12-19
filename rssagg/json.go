package main

import (
	"encoding/json"

	"log"

	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 409 {
		log.Printf("Responding with 5XX error %v", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//Return the payload to json format.
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Faimled to parse %v to json", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}