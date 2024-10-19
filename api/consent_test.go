package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConsent(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/consent", nil)
	writer := httptest.NewRecorder()
	Consent(writer, req)

	res := writer.Result()
	defer res.Body.Close()

	if resp, err := io.ReadAll(res.Body); err != nil {
		t.Fatalf("Unexpected: %s", err)
	} else {
		if string(resp) != "OK" {
			// To test output
			t.Errorf("Want OK, got %s", resp)
		}
	}
}
