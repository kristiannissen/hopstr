package api

import "net/http"

func Hop(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Hello Kitty"))
}
