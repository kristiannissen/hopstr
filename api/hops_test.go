package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHops(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hops", nil)
	w := httptest.NewRecorder()

	Hops(w, req)

	res := w.Result()
	defer res.Body.Close()

	resp, _ := io.ReadAll(res.Body)

	t.Error(string(resp))
}
