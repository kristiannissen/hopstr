package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"hopstr/api/_pkg/domain"
)

func TestHop(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hop?slug=columbia", nil)
	w := httptest.NewRecorder()

	Hop(w, req)

	res := w.Result()
	defer res.Body.Close()

	resp, _ := io.ReadAll(res.Body)

	hop := &domain.Hop{}
	json.Unmarshal(resp, &hop)

	if hop.Name != "Columbia" {
		t.Errorf("want Columbia got %s", hop.Name)
	}
}
