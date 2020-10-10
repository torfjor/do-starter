package do_starter

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type wrappedReponseWriter struct {
	status int
	n      int
	http.ResponseWriter
}

func (w *wrappedReponseWriter) Write(bytes []byte) (int, error) {
	n, err := w.ResponseWriter.Write(bytes)
	if err != nil {
		return n, err
	}
	w.n = n
	return n, nil
}

func (w *wrappedReponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func RequestLogger(out io.Writer, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		wr := &wrappedReponseWriter{200, 0, w}
		defer func(start time.Time) {
			fmt.Fprintf(out, "ts=%s method=%s url=%s status=%d len=%d took=%s\n", start.Format(time.RFC3339), r.Method, r.URL, wr.status, wr.n, time.Since(start))
		}(time.Now())
		h.ServeHTTP(wr, r)
	}

	return http.HandlerFunc(fn)
}
