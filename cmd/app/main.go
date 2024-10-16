package main

import (
	"log"
	"net/http"

	"github.com/R3iwan/blog-app/internal/blog"
	"github.com/R3iwan/blog-app/internal/db"
	"github.com/R3iwan/blog-app/internal/middleware"
	"github.com/R3iwan/blog-app/internal/user"
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
	r.HandleFunc("/api/v1/register", user.RegisterUser).Methods("POST")
	r.HandleFunc("/api/v1/login", user.LoginUser).Methods("POST")
	r.HandleFunc("/api/v1/blog/posts", blog.CreatePostHandler).Methods("POST")
	r.Handle("/api/v1/posts", middleware.JWTMiddleware(http.HandlerFunc(blog.CreatePostHandler), cfg)).Methods("POST")
	r.Handle("/api/v1/posts/{id}", middleware.JWTMiddleware(http.HandlerFunc(blog.UpdatePostHandler), cfg)).Methods("PUT")
	r.Handle("/api/v1/posts/{id}", middleware.JWTMiddleware(http.HandlerFunc(blog.DeletePostHandler), cfg)).Methods("DELETE")

	log.Printf("Server started on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
