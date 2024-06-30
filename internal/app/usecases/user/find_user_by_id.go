package user

import (
	"github.com/marcelocquadros/blog/internal/app/domain/user"
)

type (
	findUserByID struct {
		userRepository user.UserRepository
	}

	FindUserByID interface {
		Execute(id string) (*FindUserByIDResponse, error)
	}

	FindUserByIDResponse struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)

func NewGetUserByID(r user.UserRepository) *findUserByID {
	return &findUserByID{userRepository: r}
}

func (uc findUserByID) Execute(id string) (*FindUserByIDResponse, error) {
	u, err := uc.userRepository.FindUserByID(id)

	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, user.ErrUserNotFound
	}

	return &FindUserByIDResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}, nil
}
