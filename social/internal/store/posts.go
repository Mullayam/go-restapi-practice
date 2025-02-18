package store

import (
	"context"
	"database/sql"
	"errors"
)

type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}
type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	rawSql := `insert into posts (content, title, user_id, tags) values ($1, $2, $3, $4) returning id, created_at, updated_at`
	err := s.db.QueryRowContext(ctx, rawSql, post.Content, post.Title, post.UserID, post.Tags).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) GetByID(ctx context.Context, id int64) (*Post, error) {
	var post Post
	rawSql := `SELECT id, created_at, updated_at FROM posts WHERE id = $1`
	err := s.db.QueryRowContext(ctx, rawSql, int64(id)).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("post not found")
		default:
			return &post, nil
		}

	}
	return &post, nil
}
func (s *PostStore) Update(ctx context.Context, post *Post) error {
	rawSql := `UPDATE posts SET content = $1, title = $2, user_id = $3, tags = $4, updated_at = $5 WHERE id = $1`
	err := s.db.QueryRowContext(ctx, rawSql, post.ID).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) Delete(ctx context.Context, id int64) error {

	rawSql := `DELETE * FROM posts WHERE id = $1`
	err := s.db.QueryRowContext(ctx, rawSql, int64(id)).Err()
	if err != nil {
		return err
	}
	return nil
}
