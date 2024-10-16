package middleware

import (
	"context"
	"net/http"

	"github.com/R3iwan/blog-app/pkg/config"
	"github.com/golang-jwt/jwt/v4"
)

type contextKey string

const userIDKey contextKey = "userID"

func JWTMiddleware(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		cfg, err := config.NewConfig()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		tokenStr = tokenStr[7:]
		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT_Secret), nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, int((*claims)["userID"].(float64)))
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func GenerateJWT(usedID int, cfg *config.Config) (string, error) {
	claims := jwt.MapClaims{
		"userID": usedID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT_Secret))
}
