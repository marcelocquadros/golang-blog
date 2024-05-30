package post

import "github.com/marcelocquadros/blog/internal/app/domain/post"

type (
	FindAllPostsResponse struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Content  string `json:"content"`
		ImageURL string `json:"image_url"`
		UserID   string `json:"user_id"`
		Likes    int64  `json:"likes"`
		Dislikes int64  `json:"dislikes"`
	}

	findAllPosts struct {
		postRepository post.PostRepository
	}

	FindAllPosts interface {
		Execute() ([]FindAllPostsResponse, error)
	}
)

func NewFindAllPosts(r post.PostRepository) *findAllPosts {
	return &findAllPosts{postRepository: r}
}

func (f findAllPosts) Execute() ([]FindAllPostsResponse, error) {
	posts, err := f.postRepository.FindAllPosts()
	if err != nil {
		return nil, err
	}
	res := make([]FindAllPostsResponse, 0, len(posts))
	for _, p := range posts {
		res = append(res, FindAllPostsResponse{
			ID:       p.ID,
			Title:    p.Title,
			Content:  p.Content,
			ImageURL: p.ImageURL,
			UserID:   p.UserID,
			Likes:    p.Likes,
			Dislikes: p.Dislikes,
		})
	}

	return res, nil
}
