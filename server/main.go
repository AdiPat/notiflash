package main

import (
	"log"
	"main/handlers"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/notification", handlers.NotificationHandler)
	http.ListenAndServe(":8080", nil)
}
