package blog

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/R3iwan/blog-app/internal/middleware"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authorID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		log.Println("Unauthorized: missing userID in handler")
		http.Error(w, "Unauthorized: missing userID", http.StatusUnauthorized)
		return
	}

	postID, err := CreatePost(req, authorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"id": postID})
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := GetPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	var req UpdatePostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authorID := r.Context().Value(middleware.UserIDKey).(int)
	userRole := r.Context().Value(middleware.RoleKey).(string)

	if userRole != "admin" && !isPostOwner(req.ID, authorID) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	err = UpdatePost(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post updated successfully"))
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	var req DeletePostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authorID := r.Context().Value(middleware.UserIDKey).(int)
	userRole := r.Context().Value(middleware.RoleKey).(string)

	if userRole != "admin" && !isPostOwner(req.ID, authorID) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	err = DeletePost(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post deleted successfully"))
}
