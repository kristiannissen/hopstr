package api

import "net/http"

func PlaatoWebhook(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Hello Kitty"))
}
