package post

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePost_Execute(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := NewCreatePost(mockRepo)

	cmd := &CreatePostCmd{
		Title:    "Test Title",
		Content:  "Test Content",
		ImageURL: "http://example.com/image.jpg",
		UserID:   "48307347-ec81-43f0-aa2c-810d9fe8e4ee",
	}

	mockRepo.On("CreatePost", mock.AnythingOfType("*post.Post")).Return(nil)

	response, err := service.Execute(cmd)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	mockRepo.AssertExpectations(t)
}

func TestCreatePost_Execute_NewPostError(t *testing.T) {
	service := NewCreatePost(new(MockPostRepository))

	cmd := &CreatePostCmd{
		Title:    "",
		Content:  "Test Content",
		ImageURL: "http://example.com/image.jpg",
		UserID:   "48307347-ec81-43f0-aa2c-810d9fe8e4ee",
	}

	response, err := service.Execute(cmd)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestCreatePost_Execute_CreatePostError(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := NewCreatePost(mockRepo)

	cmd := &CreatePostCmd{
		Title:    "Test Title",
		Content:  "Test Content",
		ImageURL: "http://example.com/image.jpg",
		UserID:   "48307347-ec81-43f0-aa2c-810d9fe8e4ee",
	}

	mockRepo.On("CreatePost", mock.AnythingOfType("*post.Post")).Return(errors.New("create error"))

	response, err := service.Execute(cmd)
	assert.Error(t, err)
	assert.Nil(t, response)

	mockRepo.AssertExpectations(t)
}
