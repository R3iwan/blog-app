package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/R3iwan/blog-app/pkg/config"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	var req RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" || req.Email == "" || req.Role == "" {
		http.Error(w, "username, password, and email are required", http.StatusBadRequest)
	}

	err = Register(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("User %s registered", req.Username)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User successfully registered"))
}

func LoginUser(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}

	token, err := Login(req, cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	log.Printf("User %s logged in", req.Username)

	response := map[string]string{
		"access_token": token,
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User successfully logged in\n\n"))
	json.NewEncoder(w).Encode(response)
}
