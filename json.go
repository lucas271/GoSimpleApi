package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Println("Responding with 5xx error", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, statusCode, errResponse{
		Error: msg,
	})
}

func respondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {

	dat, err := json.Marshal(payload)

	if err != nil {
		log.Printf("FAILED TO MARSHAL JSON RESPONSE %s", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(dat)
}
