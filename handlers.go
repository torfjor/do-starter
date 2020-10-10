package do_starter

import (
	"context"
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}

type QuoteRetriever interface {
	GetQuote(ctx context.Context) (string, error)
}

func QuoteHandler(qr QuoteRetriever) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		quote, err := qr.GetQuote(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(quote))
	}

	return fn
}
