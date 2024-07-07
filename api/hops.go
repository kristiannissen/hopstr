package api

import (
	"encoding/json"
	"log"
	"net/http"

	"hopstr/api/_pkg/domain"
)

func Hops(w http.ResponseWriter, req *http.Request) {
	//
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	repo := domain.NewVercelHopRepository()
	hoplist := &domain.Hoplist{}
	var err error

	if hoplist, err = repo.List(); err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404"))
	} else {
		var b []byte
		if b, err = json.MarshalIndent(hoplist, "", "  "); err != nil {
			log.Println(err)

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - internal server error"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(b)
		}
	}
}
