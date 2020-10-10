package do_starter

import (
	"bytes"
	"context"
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

			gotCode := res.Code
			if gotCode != tt.wantStatus {
				t.Errorf("got status %d, want %d", tt.wantStatus, gotCode)
			}

			if !bytes.Equal(tt.want, res.Body.Bytes()) {
				t.Errorf("got body %q, want %q", res.Body.Bytes(), tt.want)
			}
		})
	}
}

type mockQuoter struct{}

func (m *mockQuoter) GetQuote(ctx context.Context) (string, error) {
	return "", nil
}

type errorQuoter struct{}

func (e *errorQuoter) GetQuote(ctx context.Context) (string, error) {
	return "", fmt.Errorf("error")
}

func TestQuoteHandler(t *testing.T) {
	tests := []struct {
		name        string
		wantContent string
		wantCode    int
		quoter      QuoteRetriever
	}{
		{
			"Responds with correct status and content-type",
			"text/plain; charset=utf-8",
			http.StatusOK,
			&mockQuoter{},
		},
		{
			"Responds with error if quoter errs",
			"text/plain; charset=utf-8",
			http.StatusInternalServerError,
			&errorQuoter{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			QuoteHandler(tt.quoter).ServeHTTP(res, req)

			gotCode := res.Code
			if gotCode != tt.wantCode {
				t.Errorf("got code %d, want %d", gotCode, tt.wantCode)
			}

			gotContent := res.Result().Header.Get("Content-Type")
			if gotContent != tt.wantContent {
				t.Errorf("got Content-Type %q, want %q", gotContent, tt.wantContent)
			}
		})
	}
}
