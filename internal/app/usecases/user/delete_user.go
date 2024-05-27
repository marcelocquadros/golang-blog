package user

import "github.com/marcelocquadros/blog/internal/app/domain/user"

type (
	deleteUser struct {
		userRepository user.UserRepository
	}

	DeleteUser interface {
		Execute(id string) error
	}
)

func NewDeleteUser(r user.UserRepository) *deleteUser {
	return &deleteUser{userRepository: r}
}

func (uc deleteUser) Execute(id string) error {
	return uc.userRepository.DeleteUser(id)
}
