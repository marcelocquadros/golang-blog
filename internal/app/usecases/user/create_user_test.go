package user

import (
	"testing"

	"github.com/marcelocquadros/blog/internal/app/domain/user"
	"github.com/marcelocquadros/blog/internal/app/usecases/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_Execute(t *testing.T) {

	mockRepo := new(mocks.MockUserRepository)
	createUserUseCase := NewCreateUser(mockRepo)

	cmd := &CreateUserCmd{
		Username: "testuser",
		Email:    "testuser@example.com",
	}

	mockRepo.On("CreateUser", mock.AnythingOfType("*user.User")).Return(nil)

	resp, err := createUserUseCase.Execute(cmd)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.ID)

}

func TestCreateUserError_Execute(t *testing.T) {

	mockRepo := new(mocks.MockUserRepository)
	createUserUseCase := NewCreateUser(mockRepo)

	cmd := &CreateUserCmd{
		Username: "testuser",
		Email:    "invalidemail",
	}

	mockRepo.On("CreateUser", mock.AnythingOfType("*user.User")).Return(nil)

	resp, err := createUserUseCase.Execute(cmd)

	assert.Empty(t, resp)
	assert.ErrorIs(t, user.ErrInvalidEmail, err)
	assert.Equal(t, err.Error(), "invalid email")

}
