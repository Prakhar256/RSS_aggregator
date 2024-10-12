package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	// jo bhi payload me data aaya usko marshal (means struct ko json format me convert) kar diya agar if no error respose writer usee writ kar dega and status code 200 for success wuld appear otherise response code for server error would be shown

	if err!=nil{
		log.Printf("Error marshalling json: %v", payload)
		w.WriteHeader(500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, statusCode int, msg string){
	if statusCode>=500{
		log.Println("Server side error", msg)
	}
	type errResponse struct{
		Error string `json:"error"`
	}
	respondWithJSON(w,statusCode,errResponse {
		Error: msg,
	})
}
