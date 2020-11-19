package routes

import (
    "net/http"
    "fmt"
    "web"
    "api"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    index := Index{Title: "Index"}

    fmt.Fprint(w, "Hello "+ index.Title)
}

func HopHandler(w http.ResponseWriter, r *http.Request) {
    hop := Hop{Name: "Amarillo"}

    fmt.Fprint(w, "Hello "+ hop.Name)
}
