package api

import "net/http"

func Hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)

	resp := make(chan []byte)

	go func() {
		resp <- []byte("Hello Pussy")
	}()

	response := <-resp

	w.Write(response)
}
