package post

import (
	"errors"
	"testing"

	"github.com/marcelocquadros/blog/internal/app/domain/post"
	"github.com/marcelocquadros/blog/internal/app/usecases/post/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindPostsByID_Execute(t *testing.T) {
	mockRepo := new(mocks.MockPostRepository)
	service := NewFindPostByID(mockRepo)

	post := &post.Post{

		ID:       "id",
		Title:    "title",
		Content:  "content",
		ImageURL: "image_url",
		UserID:   "userid",
		Likes:    30,
		Dislikes: 1,
	}

	mockRepo.On("FindPostByID", mock.AnythingOfType("string")).Return(post, nil)
	res, err := service.Execute("id")
	assert.NoError(t, err)
	assert.Equal(t, res.ID, "id")
	assert.Equal(t, res.Title, "title")
	assert.Equal(t, res.ImageURL, "image_url")
	assert.Equal(t, res.UserID, "userid")
	assert.EqualValues(t, res.Likes, 30)
	assert.EqualValues(t, res.Dislikes, 1)
	mockRepo.AssertExpectations(t)
}

func TestFindPostByID_ExecuteError(t *testing.T) {
	mockRepo := new(mocks.MockPostRepository)
	service := NewFindPostByID(mockRepo)
	mockRepo.On("FindPostByID", mock.AnythingOfType("string")).Return(nil, errors.New("error"))
	_, err := service.Execute("ids")
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFindPostByID_ExecuteErrorNotFound(t *testing.T) {
	mockRepo := new(mocks.MockPostRepository)
	service := NewFindPostByID(mockRepo)
	mockRepo.On("FindPostByID", mock.AnythingOfType("string")).Return(nil, post.ErrPostNotFound)
	_, err := service.Execute("id")
	assert.Error(t, err)
	assert.ErrorIs(t, err, post.ErrPostNotFound)
	mockRepo.AssertExpectations(t)
}
