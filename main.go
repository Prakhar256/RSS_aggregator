package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	//this will fetch all environment variables form my .env file and put it into my current environment

	PortString := os.Getenv("PORT")
	if PortString == "" {
		log.Fatal(PortString + "is not available in environment")
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handleReadiness)
	v1Router.Get("/err", handleErrors)
	router.Mount("/v1", v1Router)

	port := ":" + PortString
	fmt.Println("Starting server on port", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
