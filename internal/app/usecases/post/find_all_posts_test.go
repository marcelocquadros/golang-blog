package post

import (
	"errors"
	"testing"

	"github.com/marcelocquadros/blog/internal/app/domain/post"
	"github.com/stretchr/testify/assert"
)

func TestFindAllPosts_Execute(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := NewFindAllPosts(mockRepo)

	posts := []post.Post{
		{
			ID:       "id",
			Title:    "title",
			Content:  "content",
			ImageURL: "image_url",
			UserID:   "userid",
			Likes:    30,
			Dislikes: 1,
		},
	}

	mockRepo.On("FindAllPosts").Return(posts, nil)
	res, err := service.Execute()
	assert.NoError(t, err)
	assert.Equal(t, res[0].ID, "id")
	assert.Equal(t, res[0].Title, "title")
	assert.Equal(t, res[0].ImageURL, "image_url")
	assert.Equal(t, res[0].UserID, "userid")
	assert.EqualValues(t, res[0].Likes, 30)
	assert.EqualValues(t, res[0].Dislikes, 1)
	mockRepo.AssertExpectations(t)
}

func TestFindAllPosts_ExecuteError(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := NewFindAllPosts(mockRepo)

	mockRepo.On("FindAllPosts").Return([]post.Post{}, errors.New("error"))
	_, err := service.Execute()
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
