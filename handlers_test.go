package do_starter

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name       string
		param      string
		want       []byte
		wantStatus int
	}{
		{"Greets name", "John", []byte("Hello, John!"), http.StatusOK},
		{"Greets world when no name is given", "", []byte("Hello, world!"), http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := httptest.NewRecorder()
			req, _ := http.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/?name=%s", tt.param), nil)
			HelloHandler(res, req)

			gotCode := res.Result().StatusCode
			if gotCode != tt.wantStatus {
				t.Errorf("got status %d, want %d", tt.wantStatus, gotCode)
			}

			if !bytes.Equal(tt.want, res.Body.Bytes()) {
				t.Errorf("got body %q, want %q", res.Body.Bytes(), tt.want)
			}
		})
	}
}
