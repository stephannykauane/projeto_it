package middleware

import (
	"context"
	"go/types"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stephannykauane/projeto_it/backend/headers"
)


type contextKey string 
const userContextKey = contextKey("userEmail")
var jwtKey = os.Getenv("JWT_SECRET")


func Auth(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return 
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// AQUI É A MUDANÇA: usar types.Claims
		token, err := jwt.ParseWithClaims(tokenStr, &types.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token Inválido", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(*types.Claims)
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


