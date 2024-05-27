package user

import "github.com/marcelocquadros/blog/internal/app/domain/user"

type (
	updateUser struct {
		userRepository user.UserRepository
	}

	UpdateUserCmd struct {
		Email    *string `json:"email" binding:"email"`
		Username *string `json:"username"`
	}

	UpdateUser interface {
		Execute(cmd UpdateUserCmd) (*UserResponse, error)
	}
)

func NewUpdateUser(r user.UserRepository) *updateUser {
	return &updateUser{userRepository: r}
}

func (uc updateUser) Execute(cmd UpdateUserCmd) (*UserResponse, error) {
	return nil, nil
}
