package middleware

import (
    
    "log"
    "net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
    for _, m := range middlewares {
        f = m(f)
    }
    return f
}

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next(w, r)
    }
}

func ValidateJSON(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("Content-Type") != "application/json" {
            http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
            return
        }
        next(w, r)
    }
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        apiKey := r.Header.Get("X-API-Key")
        if apiKey == "" {
            http.Error(w, "Missing API key", http.StatusUnauthorized)
            return
        }
        // In a real application, you would validate the API key against a database or service
        if apiKey != "your-api-key" {
            http.Error(w, "Invalid API key", http.StatusUnauthorized)
            return
        }
        next(w, r)
    }
}