package post

import (
	"errors"
	"testing"

	"github.com/marcelocquadros/blog/internal/app/usecases/post/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeletePost_Execute(t *testing.T) {
	mockRepo := new(mocks.MockPostRepository)
	service := NewDeletePost(mockRepo)

	mockRepo.On("DeletePost", mock.AnythingOfType("string")).Return(nil)
	id := "d72cf453-1f6a-4474-9968-a45c7560bf4c"
	err := service.Execute(id)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeletePost_ExecuteError(t *testing.T) {
	mockRepo := new(mocks.MockPostRepository)
	service := NewDeletePost(mockRepo)

	mockRepo.On("DeletePost", mock.AnythingOfType("string")).Return(errors.New("DeletePost error"))
	id := "d72cf453-1f6a-4474-9968-a45c7560bf4c"
	err := service.Execute(id)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
