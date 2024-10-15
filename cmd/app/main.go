package main

import (
	"log"
	"net/http"

	"github.com/R3iwan/blog-app/internal/db"
	"github.com/R3iwan/blog-app/pkg/config"
	"github.com/R3iwan/blog-app/pkg/logger"
	"github.com/gorilla/mux"
)

func init() {
	logger.InitLogger()
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	connDB, err := db.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer connDB.Close()

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/register", RegisterUser).Methods("POST")
	r.HandleFunc("/api/v1/login", LoginUser).Methods("POST")

	log.Printf("Server started on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Print("User registered")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User registered"))
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	log.Print("User logged in")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User logged in"))
}
