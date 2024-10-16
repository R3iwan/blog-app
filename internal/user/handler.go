package user

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/R3iwan/blog-app/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req User

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" || req.Email == "" {
		http.Error(w, "username, password, and email are required", http.StatusBadRequest)
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Password = hashedPassword

	tx, err := db.DB.Begin(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec(context.Background(), "INSERT INTO users (username, password_hash, email) VALUES ($1, $2, $3)", req.Username, req.Password, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("User %s registered", req.Username)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User successfully registered"))
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var req User

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}

	hashedPassword, err := getHashPassword(req.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}

	log.Printf("User %s logged in", req.Username)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User successfully logged in"))
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
