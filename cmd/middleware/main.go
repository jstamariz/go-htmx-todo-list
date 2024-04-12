package middleware

import (
	"net/http"
	"regexp"

	"github.com/jstamariz/go-htmx/cmd/logger"
)

func XSSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		xssPattern := regexp.MustCompile(`(<|>|&lt;|&gt;|alert\(|script\(|javascript:|on\w+=)`)

		// Check each form parameter for XSS patterns
		for _, values := range r.Form {
			for _, value := range values {
				if xssPattern.MatchString(value) {
					http.Error(w, "Potential XSS attack detected", http.StatusForbidden)
					logger.Log.Printf("Potential XSS attack detected at %s", r.URL)
					return
				}
			}
		}

		next(w, r)
	}
}
