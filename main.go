package main

import (
	"log"
	"net/http"
	"os"
  "fmt"
	"thingy"
  "routes"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  substitutes := []thingy.Thingy{
    thingy.Thingy{Name: "hops"},
  }
  thingy := thingy.Thingy{Name: "Pussy", Substitutes: substitutes}
  fmt.Fprintf(w, "Hello %s", thingy.Name)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("Default port %s", port)
	}

  http.HandleFunc("/", routes.IndexHandler)
  http.HandleFunc("/hops/", routes.HopsHandler)

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
