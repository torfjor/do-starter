package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	do_starter "github.com/torfjor/do-starter"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/greet", do_starter.HelloHandler)
	mux.HandleFunc("/quote", do_starter.QuoteHandler(&do_starter.RandomQuoter{}))

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      do_starter.RequestLogger(os.Stderr, mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ListenAndServe: %v", err)
		os.Exit(1)
	}
}
