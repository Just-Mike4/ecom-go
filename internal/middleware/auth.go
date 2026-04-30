package middleware

import (
	"net/http"

	"github.com/Just-Mike4/go-ecom/internal/json"
)

// AuthMiddleware validates JWT tokens
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement JWT validation
		// For now, just pass through
		next.ServeHTTP(w, r)
	})
}

// AdminMiddleware checks if user is admin
func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Check user role from JWT claims
		// For now, just pass through
		next.ServeHTTP(w, r)
	})
}

// VendorMiddleware checks if user is vendor
func VendorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Check user role from JWT claims
		// For now, just pass through
		next.ServeHTTP(w, r)
	})
}

// RateLimitMiddleware limits request rates
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement rate limiting
		next.ServeHTTP(w, r)
	})
}

// CORSMiddleware handles CORS
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// RequireAuth middleware that returns 401 if not authenticated
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Get token from Authorization header and validate
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			json.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
			return
		}

		// For now, just pass through
		// In production, validate JWT and extract user info
		next.ServeHTTP(w, r)
	})
}
