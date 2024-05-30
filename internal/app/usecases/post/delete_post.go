package post

import "github.com/marcelocquadros/blog/internal/app/domain/post"

type (
	deletePost struct {
		postRepository post.PostRepository
	}

	DeletePost interface {
		Execute(string) error
	}
)

func NewDeletePost(p post.PostRepository) *deletePost {
	return &deletePost{postRepository: p}
}

func (d deletePost) Execute(id string) error {
	if err := d.postRepository.DeletePost(id); err != nil {
		return err
	}
	return nil
}
