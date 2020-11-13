package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func main() {
  http.HandleFunc("/", indexHandler)
  
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("defaulting to port %s", port)
  }

  log.Printf("listening on port %s", port)
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal(err)
  }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  fmt.Fprint(w, "Hello Kitty")
}
