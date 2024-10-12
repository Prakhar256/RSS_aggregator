package main

import "net/http"
func handleErrors(w http.ResponseWriter, r *http.Request){
	respondWithError(w,400,"Something went wrong")
}