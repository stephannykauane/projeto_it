package headers

import (
	"net/http"
    
)

func SetHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        w.Header().Set("Access-Control-Allow-Credentials", "true")

    
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

        next.ServeHTTP(w, r)
    })
}

