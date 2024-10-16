package blog

import (
	"context"

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
	rows, err := db.DB.Query(context.Background(), "SELECT id, title, content, author_id, created_at FROM posts")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt)
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
