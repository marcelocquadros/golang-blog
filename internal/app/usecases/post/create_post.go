package post

import "github.com/marcelocquadros/blog/internal/app/domain/post"

type (
	createPost struct {
		postRepository post.PostRepository
	}

	CreatePostCmd struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		ImageURL string `json:"image_url"`
		UserID   string `json:"user_id"`
	}

	CreatePostResponse struct {
		ID string `json:"id"`
	}

	CreatePost interface {
		Execute(*CreatePostCmd) (*CreatePostResponse, error)
	}
)

func NewCreatePost(r post.PostRepository) *createPost {
	return &createPost{postRepository: r}
}

func (c createPost) Execute(cmd *CreatePostCmd) (*CreatePostResponse, error) {
	p, err := post.NewPost(cmd.UserID, cmd.Title, cmd.Content, cmd.ImageURL)

	if err != nil {
		return nil, err
	}

	err = c.postRepository.CreatePost(p)
	if err != nil {
		return nil, err
	}

	return &CreatePostResponse{ID: p.ID}, nil
}
