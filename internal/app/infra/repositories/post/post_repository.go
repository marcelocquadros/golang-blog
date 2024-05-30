package post

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/marcelocquadros/blog/internal/app/domain/post"
)

type (
	postRepository struct {
		db *sqlx.DB
	}
)

func NewPostRepository(db *sqlx.DB) *postRepository {
	return &postRepository{db: db}
}

func (r postRepository) CreatePost(p *post.Post) error {
	if _, err := r.db.Exec("INSERT INTO posts (id, title, content, image_url, user_id) VALUES(?, ?, ?, ?, ?)", p.ID, p.Title, p.Content, p.ImageURL, p.UserID); err != nil {
		return err
	}
	return nil
}

func (r postRepository) UpdatePost(p *post.Post) error {
	if _, err := r.db.Exec("UPDATE posts SET title=?, content=?, image_url=? WHERE id=?", p.Title, p.Content, p.ImageURL, p.ID); err != nil {
		return err
	}
	return nil
}

func (r postRepository) DeletePost(id string) error {
	if _, err := r.db.Exec("DELETE FROM posts WHERE id=?", id); err != nil {
		return err
	}
	return nil
}

func (r postRepository) FindAllPosts() ([]post.Post, error) {
	posts := make([]post.Post, 0)
	if err := r.db.Select(&posts, "SELECT * FROM posts"); err != nil {
		return []post.Post{}, err
	}
	return posts, nil
}

func (r postRepository) FindPostByID(id string) (*post.Post, error) {
	post := post.Post{}
	if err := r.db.Get(&post, "SELECT * FROM posts WHERE id=?", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}
