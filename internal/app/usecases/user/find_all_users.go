package user

import "github.com/marcelocquadros/blog/internal/app/domain/user"

type (
	FindAllUsersResponse struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	findAllUsers struct {
		userRepository user.UserRepository
	}

	FindAllUsers interface {
		Execute() ([]FindAllUsersResponse, error)
	}
)

func NewFindAllUsers(r user.UserRepository) *findAllUsers {
	return &findAllUsers{userRepository: r}
}

func (uc findAllUsers) Execute() ([]FindAllUsersResponse, error) {
	users, err := uc.userRepository.FindAllUsers()

	if err != nil {
		return nil, err
	}

	res := make([]FindAllUsersResponse, 0, len(users))
	for _, u := range users {
		res = append(res, FindAllUsersResponse{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		})
	}

	return res, nil
}
