package main

import (
	"log"
	"net/http"
)

func main() {
	// Set up an HTTP server to receive the webhook events
	http.HandleFunc("/webhook", HandleWebhook)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
