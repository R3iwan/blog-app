package blog

import (
	"context"
	"log"

	"github.com/R3iwan/blog-app/internal/db"
)

func CreatePost(req CreatePostRequest, author_id int) (int, error) {
	var postID int
	query := `INSERT INTO posts(title, content, author_id) VALUES($1, $2, $3) RETURNING id`

	err := db.DB.QueryRow(context.Background(), query, req.Title, req.Content, author_id).Scan(&postID)
	if err != nil {
		return 0, err
	}
	return postID, nil
}

func GetPosts() ([]Post, error) {
	var posts []Post
	rows, err := db.DB.Query(context.Background(),
		`SELECT p.id, p.title, p.content, u.username, p.created_at
	FROM posts p 
	JOIN users u ON p.author_id = u.id`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Username, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func UpdatePost(req UpdatePostRequest) error {
	query := `UPDATE posts SET title=$1, content=$2 WHERE id=$3`

	_, err := db.DB.Exec(context.Background(), query, req.Title, req.Content, req.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(req DeletePostRequest) error {
	query := `DELETE FROM posts WHERE id=$1`

	_, err := db.DB.Exec(context.Background(), query, req.ID)
	if err != nil {
		return err
	}
	return nil
}

func isPostOwner(postID, userID int) bool {
	var ownerID int
	err := db.DB.QueryRow(context.Background(), "SELECT author_id FROM posts WHERE id = $1", postID).Scan(&ownerID)
	if err != nil {
		log.Printf("Error checking post ownership: %v", err)
		return false
	}

	return ownerID == userID
}
