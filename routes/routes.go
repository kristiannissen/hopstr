package routes

import (
    "net/http"
    "fmt"
    "web"
    "api"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    assets := make(map[string]string)
    assets["css"] = "index.css"
    assets["js"] = "index.js"

    index := Index{Title: "Index", Assets: assets}

    fmt.Fprint(w, "Hello "+ index.Title)
}

func HopHandler(w http.ResponseWriter, r *http.Request) {
    hop := Hop{Name: "Amarillo"}

    fmt.Fprint(w, "Hello "+ hop.Name)
}
