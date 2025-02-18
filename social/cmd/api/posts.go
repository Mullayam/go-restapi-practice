package main

import (
	"net/http"

	"github.com/mullayam/social/internal/store"
)

type CreatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var post *store.Post
	if err := readJSON(w, r, post); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	userId := 1
	post = &store.Post{Title: post.Title, Content: post.Content, UserID: int64(userId)}
	if err := app.store.Posts.Create(r.Context(), post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
