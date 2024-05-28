package user

import "github.com/marcelocquadros/blog/internal/app/domain/user"

type (
	createUser struct {
		userRepository user.UserRepository
	}

	CreateUser interface {
		Execute(cmd *CreateUserCmd) (CreateUserResponse, error)
	}

	CreateUserCmd struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}

	CreateUserResponse struct {
		ID string `json:"id"`
	}
)

func NewCreateUser(r user.UserRepository) *createUser {
	return &createUser{userRepository: r}
}

func (uc *createUser) Execute(cmd *CreateUserCmd) (*CreateUserResponse, error) {
	user, err := user.NewUser(cmd.Username, cmd.Email)

	if err != nil {
		return nil, err
	}

	if err := uc.userRepository.CreateUser(user); err != nil {
		return nil, err
	}

	return &CreateUserResponse{ID: user.ID}, nil
}
