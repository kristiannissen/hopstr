package main

import (
	"log"
	"net/http"

	a "hopstr/api"
)

func main() {
	mux := http.NewServeMux()
	// Dummy handler
	mux.Handle("/api/consent", http.HandlerFunc(a.Consent))
	// Hop handler
	mux.Handle("/api/hop", http.HandlerFunc(a.Hop))
	// Hops handler
	mux.Handle("/api/hops", http.HandlerFunc(a.Hops))
	// Plaato webhook
	mux.Handle("/api/plaatowebhook", http.HandlerFunc(a.PlaatoWebhook))

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":1234", mux))
}
