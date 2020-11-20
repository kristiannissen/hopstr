package routes

import (
    "net/http"
    "fmt"
    "routes/web"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    assets := map[string]string{"css": "css.css", "js": "js.js"}

    index := web.Index{Title: "Index", Assets: assets}

    fmt.Fprint(w, "Hello "+ index.Title)

    for k, v := range index.Assets {
      fmt.Fprint(w, "Asset: "+ k +":"+ v)
    }
}
