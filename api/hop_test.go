package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHop(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hop", nil)
	w := httptest.NewRecorder()

	Hop(w, req)

	res := w.Result()
	defer res.Body.Close()

	resp, _ := io.ReadAll(res.Body)

	t.Error(string(resp))
}
