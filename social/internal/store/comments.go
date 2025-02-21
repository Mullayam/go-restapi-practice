package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	UserID    int64  `json:"user_id"`
	PostID    int64  `json:"post_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type CommentStore struct {
	db *sql.DB
}

func NewCommentStore(db *sql.DB) CommentStore {
	return CommentStore{db}
}
func (s *CommentStore) Create(ctx context.Context, comment *Comment) error {

}
