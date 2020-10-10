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

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      http.HandlerFunc(do_starter.HelloHandler),
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
