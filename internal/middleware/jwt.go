package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/R3iwan/blog-app/pkg/config"
	"github.com/golang-jwt/jwt/v4"
)

type contextKey string

const (
	UserIDKey contextKey = "userID"
	RoleKey   contextKey = "role"
)

func JWTMiddleware(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr = tokenStr[7:]
		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT_Secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID, ok := (*claims)["userID"].(float64)
		if !ok {
			http.Error(w, "Invalid token data", http.StatusUnauthorized)
			return
		}

		role, ok := (*claims)["role"].(string)
		if !ok {
			http.Error(w, "Invalid token data", http.StatusUnauthorized)
			return
		}

		userIDVal, ok := (*claims)["userID"]
		if !ok || userIDVal == nil {
			log.Println("userID missing from token")
			http.Error(w, "Unauthorized: missing userID", http.StatusUnauthorized)
			return
		}
		
		ctx := context.WithValue(r.Context(), UserIDKey, int(userID))
		ctx = context.WithValue(ctx, RoleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func GenerateJWT(userID int, username string, role string, duration time.Time, cfg *config.Config) (string, error) {
	claims := jwt.MapClaims{
		"userID":   userID,
		"username": username,
		"role":     role,
		"exp":      duration.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT_Secret))
}
