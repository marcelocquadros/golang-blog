package post

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrInvalidPostID   = errors.New("invalid id format")
	ErrInvalidTitle    = errors.New("invalid title")
	ErrInvalidContent  = errors.New("invalid content")
	ErrInvalidImageURL = errors.New("invalid image url")
	ErrInvalidUserID   = errors.New("invalid user id")
	ErrPostNotFound    = errors.New("post not found")
)

type (
	Post struct {
		ID       string
		Title    string
		Content  string
		ImageURL string `db:"image_url"`
		UserID   string `db:"user_id"`
		Likes    int64
		Dislikes int64
	}

	PostRepository interface {
		CreatePost(p *Post) error
		UpdatePost(p *Post) error
		DeletePost(id string) error
		FindAllPosts() ([]Post, error)
		FindPostByID(id string) (*Post, error)
	}
)

func NewPost(userID string, title string, content string, imageURL string) (*Post, error) {
	p := &Post{
		ID:       uuid.NewString(),
		Title:    title,
		UserID:   userID,
		Content:  content,
		ImageURL: imageURL,
	}
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p Post) IsValid() (bool, error) {
	if _, err := uuid.Parse(p.ID); err != nil {
		return false, ErrInvalidPostID
	}
	if isEmpty(p.Title) {
		return false, ErrInvalidTitle
	}

	if isEmpty(p.Content) {
		return false, ErrInvalidContent
	}

	if !isValidURL(p.ImageURL) {
		return false, ErrInvalidImageURL
	}

	if _, err := uuid.Parse(p.UserID); err != nil {
		return false, ErrInvalidUserID
	}

	return true, nil
}

func isEmpty(s string) bool {
	return len(s) == 0
}

func isValidURL(u string) bool {
	if !isEmpty(u) {
		if strings.HasPrefix(u, "https://") || strings.HasPrefix(u, "http://") {
			return true
		}
	}
	return false
}
