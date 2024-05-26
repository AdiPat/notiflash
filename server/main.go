package main

import (
	"main/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/notification", handlers.NotificationHandler)
	http.ListenAndServe(":8080", nil)
}
