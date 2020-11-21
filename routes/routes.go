package routes

import (
    "net/http"
    "fmt"
    "routes/web"
    "routes/api"
  )

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    assets := map[string]string{"css": "css.css", "js": "js.js"}

    index := web.Index{Title: "Index", Assets: assets}

    fmt.Fprint(w, "Hello "+ index.Title)

    for k, v := range index.Assets {
      fmt.Fprint(w, "Asset: "+ k +":"+ v)
    }
}

func HopsHandler(w http.ResponseWriter, r *http.Request) {
    hops, status := api.LoadHops()
    if status == 1 {
        fmt.Fprint(w, "Hello "+ hops["1"])
    } else {
        fmt.Fprint(w, "Hello "+ hops["2"])
    }
}
