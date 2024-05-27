package user

import "github.com/marcelocquadros/blog/internal/app/domain/user"

type (
	getAllUsers struct {
		userRepository user.UserRepository
	}

	GetAllUsers interface {
		Execute() ([]UserResponse, error)
	}
)

func NewGetAllUsers(r user.UserRepository) *getAllUsers {
	return &getAllUsers{userRepository: r}
}

func (uc getAllUsers) Execute() ([]UserResponse, error) {
	users, err := uc.userRepository.FindAllUsers()

	if err != nil {
		return nil, err
	}

	res := make([]UserResponse, len(users))
	for _, u := range users {
		res = append(res, UserResponse{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		})
	}

	return res, nil

}
