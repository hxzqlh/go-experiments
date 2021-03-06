package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(2, 5)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)
	// Wrap the servemux with the limit middleware.
	http.ListenAndServe(":4000", limit(mux))
}
