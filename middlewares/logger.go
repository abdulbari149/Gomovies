package middlewares

import (
	"fmt"
	"net/http"
)

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Do stuff here
		fmt.Printf("Request received URL=%s | Method=%s\n", r.URL, r.Method)

		// Call the next handler
		h.ServeHTTP(w, r)

		// Do stuff here
		fmt.Printf("Response sent URL=%s | Method=%s\n", r.URL, r.Method)
	})
}
