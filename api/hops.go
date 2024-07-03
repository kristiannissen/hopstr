package api

import (
	"net/http"
)

func Hops(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Have some hops"))
}
