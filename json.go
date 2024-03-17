package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to parse JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	if code > 499 {
		log.Printf("Error: %v", msg)
		w.WriteHeader(500)

		respondWithJSON(w, 500, errorResponse{
			Error: "internal server error",
		})
		return
	}
	respondWithJSON(w, 500, errorResponse{
		Error: msg,
	})

}
