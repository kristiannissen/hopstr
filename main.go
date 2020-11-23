package main

import (
	"log"
	"net/http"
	"os"
	"routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("Default port %s", port)
	}

	http.HandleFunc("/", routes.IndexHandler)
	http.HandleFunc("/api/hops/", routes.HopsHandler)
	http.HandleFunc("/api/hop/", routes.HopHandler)

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
