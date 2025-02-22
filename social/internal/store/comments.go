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
func (s *CommentStore) GetByPostID(ctx context.Context, comment *Comment) ([]*Comment, error) {
	query := `SELECT id, created_at, updated_at FROM comments WHERE post_id = $1`
	rows, err := s.db.QueryContext(ctx, query, comment.PostID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comments := make([]*Comment, 0)
	for rows.Next() {
		var c Comment
		err := rows.Scan(&comment.ID, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &c)
	}
	return comments, nil

}
