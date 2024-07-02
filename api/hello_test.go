package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	writer := httptest.NewRecorder()

	Hello(writer, req)

	res := writer.Result()
	defer res.Body.Close()

	if resp, err := io.ReadAll(res.Body); err != nil {
		t.Fatalf("Unexpected: %s", err)
	} else {
		if string(resp) != "Hello Kitty" {
			t.Errorf("Want Hello Kitty, got %s", resp)
		}
	}
}
