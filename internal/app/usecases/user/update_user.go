package user

import "github.com/marcelocquadros/blog/internal/app/domain/user"

type (
	updateUser struct {
		userRepository user.UserRepository
	}

	UpdateUser interface {
		Execute(cmd UpdateUserCmd) (*UpdateUserResponse, error)
	}

	UpdateUserResponse struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	UpdateUserCmd struct {
		Email    *string `json:"email" binding:"email"`
		Username *string `json:"username"`
	}
)

func NewUpdateUser(r user.UserRepository) *updateUser {
	return &updateUser{userRepository: r}
}

func (uc updateUser) Execute(cmd UpdateUserCmd) (*UpdateUserResponse, error) {
	return nil, nil
}
