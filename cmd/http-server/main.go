package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	do_starter "github.com/torfjor/do-starter"
	"github.com/torfjor/do-starter/pkg/chucknorris"
	"github.com/torfjor/do-starter/pkg/trump"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/greet", do_starter.HelloHandler)
	mux.HandleFunc("/trump", do_starter.QuoteHandler(&trump.Quoter{}))
	mux.HandleFunc("/chucknorris", do_starter.QuoteHandler(&chucknorris.Quoter{}))

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      do_starter.RequestLogger(os.Stderr, mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ListenAndServe: %v", err)
		os.Exit(1)
	}
}
