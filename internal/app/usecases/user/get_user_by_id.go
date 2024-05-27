package user

import (
	"github.com/marcelocquadros/blog/internal/app/domain/user"
)

type (
	getUserByID struct {
		userRepository user.UserRepository
	}

	UserResponse struct {
		ID       string
		Username string
		Email    string
	}

	GetUserByID interface {
		Execute(id string) (*UserResponse, error)
	}
)

func NewGetUserByID(r user.UserRepository) *getUserByID {
	return &getUserByID{userRepository: r}
}

func (uc getUserByID) Execute(id string) (*UserResponse, error) {
	u, err := uc.userRepository.FindUserByID(id)

	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}, nil
}
