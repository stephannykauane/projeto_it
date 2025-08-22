package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stephannykauane/projeto_it/backend/types"
)

type contextKey string

const userContextKey = contextKey("userEmail")

func Auth(next http.HandlerFunc) http.Handler {

jwtKey := os.Getenv("JWT_SECRET")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &types.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				http.Error(w, "Método de assinatura inválido", http.StatusUnauthorized)
				return nil, nil
			}
			return []byte(jwtKey), nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors&jwt.ValidationErrorExpired != 0 {
				http.Error(w, "Sessão expirada", http.StatusExpectationFailed) 
				return
			}
			http.Error(w, "Token Inválido", http.StatusExpectationFailed)
			return
		}

		if !token.Valid {
			http.Error(w, "Token Inválido", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserEmailFromContext(ctx context.Context) string {
	if email, ok := ctx.Value(userContextKey).(string); ok {
		return email
	}
	return ""
}
