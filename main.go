package main

import (
	"log"
	"net/http"
	"os"
    "fmt"
    "regexp"
)

func main() {

    port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("Default port %s", port)
	}

    routes := make(map[string]string)
    routes["^\\/$"] = "hello kitty"
    routes["\\/hops\\/$"] = "hello pussy"
    routes["\\/hops/(\\w+)\\/$"] = "hello kitty pussy"

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var path string = r.URL.Path
        for k, v := range routes {
            m, err := regexp.MatchString(k, path)
            fmt.Println(m, err, v)
        }
        fmt.Fprint(w, "Done "+ path)
    })

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
