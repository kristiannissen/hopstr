package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"hopstr/api/_pkg/domain"
)

func TestHops(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hops", nil)
	w := httptest.NewRecorder()

	Hops(w, req)

	res := w.Result()
	defer res.Body.Close()

	resp, _ := io.ReadAll(res.Body)
	hoplist := &domain.Hoplist{}

	if err := json.Unmarshal(resp, &hoplist); err != nil {
		t.Error(err)
	} else {
		if len(hoplist.Hoplist) == 0 {
			t.Error("List is empty")
		}
	}
}
