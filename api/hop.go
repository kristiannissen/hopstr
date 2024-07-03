package api

import (
	"encoding/json"
	"log"
	"net/http"

	"hopstr/api/_pkg/domain"
)

func Hop(w http.ResponseWriter, req *http.Request) {
	//
	w.Header().Set("Content-type", "application/json; charset=utf-8")

	//
	repo := domain.NewVercelHopRepository()
	hop := &domain.Hop{}
	var err error
	slug := req.URL.Query().Get("slug")

	if hop, err = repo.Find(slug); err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - not found"))
	} else {
		var b []byte
		if b, err = json.MarshalIndent(hop, "", "\t"); err != nil {
			// Internal error
			log.Fatal(err)

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - internal server error"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(b)
		}
	}
}
