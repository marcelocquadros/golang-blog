package user

import "github.com/marcelocquadros/blog/internal/app/domain/user"

type (
	createUser struct {
		userRepository user.UserRepository
	}

	CreateUser interface {
		Execute(cmd *CreateUserCmd) (string, error)
	}

	CreateUserCmd struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}
)

func NewCreateUser(r user.UserRepository) *createUser {
	return &createUser{userRepository: r}
}

func (uc *createUser) Execute(cmd *CreateUserCmd) (string, error) {
	user, err := user.NewUser(cmd.Username, cmd.Email)

	if err != nil {
		return "", err
	}

	if err := uc.userRepository.CreateUser(user); err != nil {
		return "", err
	}

	return user.ID, nil
}
