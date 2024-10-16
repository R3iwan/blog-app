package user

import (
	"context"

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

	_, err = db.DB.Exec(context.Background(), "INSERT INTO users (username, password_hash, email) VALUES ($1, $2, $3)", req.Username, hashedPassword, req.Email)
	return err
}

func Login(req LoginRequest) (string, error) {
	hashedPassword, userID, err := getHashPassword(req.Username)
	if err != nil {
		return "", err
	}

	cfg, err := config.NewConfig()
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return "", err
	}

	token, jwtErr := middleware.GenerateJWT(userID, cfg)
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

func getHashPassword(username string) (string, int, error) {
	var password string
	var userID int

	err := db.DB.QueryRow(context.Background(), "SELECT password_hash FROM users WHERE username = $1", username).Scan(&password)
	if err != nil {
		return "", 0, err
	}

	return password, userID, nil
}
