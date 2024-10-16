package user

import (
	"context"

	"github.com/R3iwan/blog-app/internal/db"
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

func Login(req LoginRequest) error {
	hashedPassword, err := getHashPassword(req.Username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	return err
}

func hashPassword(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bcryptPassword), nil
}

func getHashPassword(username string) (string, error) {
	var password string

	err := db.DB.QueryRow(context.Background(), "SELECT password_hash FROM users WHERE username = $1", username).Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil
}
