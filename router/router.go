package router

import (
	"fmt"
	"net/http"
)

func AddRoute(path string, r string) {
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello from Index")
}
