package post

import "github.com/marcelocquadros/blog/internal/app/domain/post"

type (
	findPostByID struct {
		postRepository post.PostRepository
	}

	FindPostByID interface {
		Execute(string) (*FindPostByIDResponse, error)
	}

	FindPostByIDResponse struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Content  string `json:"content"`
		ImageURL string `json:"image_url"`
		UserID   string `json:"user_id"`
		Likes    int64  `json:"likes"`
		Dislikes int64  `json:"dislikes"`
	}
)

func NewFindPostByID(r post.PostRepository) *findPostByID {
	return &findPostByID{postRepository: r}
}

func (f findPostByID) Execute(id string) (*FindPostByIDResponse, error) {
	p, err := f.postRepository.FindPostByID(id)

	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, post.ErrPostNotFound
	}

	return &FindPostByIDResponse{
		ID:       p.ID,
		Title:    p.Title,
		Content:  p.Content,
		ImageURL: p.ImageURL,
		UserID:   p.UserID,
		Likes:    p.Likes,
		Dislikes: p.Dislikes,
	}, nil
}
