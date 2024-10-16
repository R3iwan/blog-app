package user

import (
	"context"
	"fmt"
	"time"

	"github.com/R3iwan/blog-app/internal/db"
	"github.com/R3iwan/blog-app/internal/middleware"
	"github.com/R3iwan/blog-app/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

func Register(req RegisterRequest) error {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec(context.Background(), "INSERT INTO users (username, password_hash, email, role) VALUES ($1, $2, $3, $4)", req.Username, hashedPassword, req.Email, req.Role)
	return err
}

func Login(req LoginRequest, cfg *config.Config) (string, error) {
	hashedPassword, userID, dbRole, err := getUserDetails(req.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return "", err
	}

	if req.Role != dbRole {
		return "", fmt.Errorf("role mismatch")
	}

	token, jwtErr := middleware.GenerateJWT(userID, req.Username, dbRole, time.Now().Add(15*time.Minute), cfg)
	if jwtErr != nil {
		return "", jwtErr
	}

	return token, nil
}

func hashPassword(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bcryptPassword), nil
}

func getUserDetails(username string) (string, int, string, error) {
	var password string
	var userID int
	var role string

	err := db.DB.QueryRow(context.Background(), "SELECT password_hash, id, role FROM users WHERE username = $1", username).Scan(&password, &userID, &role)
	if err != nil {
		return "", 0, "", err
	}

	return password, userID, role, nil
}
