package main

import (
	"log"
	"net/http"
	"os"
  d "dispatch"
  "fmt"
)

func HelloKitty(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Hello Kitty")
}

func HelloPussy(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Hello Pussy")
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("Default port %s", port)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var path string = r.URL.Path

    disp := d.NewRoute()
    disp.AddRoute("/kitty/{hello}/", HelloKitty)
    disp.AddRoute("/pussy/", HelloPussy)
    disp.Dispatch(path)
	})

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
