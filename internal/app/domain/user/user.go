package user

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

var (
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidUserID   = errors.New("invalid id format")
)

type (
	User struct {
		ID       string
		Username string
		Email    string
	}

	UserRepository interface {
		CreateUser(u *User) error
		DeleteUser(ID string) error
		UpdateUser(u *User) error
		FindAllUsers() ([]User, error)
		FindUserByID(ID string) (*User, error)
	}
)

func NewUser(username string, email string) (*User, error) {

	u := &User{
		ID:       uuid.NewString(),
		Username: username,
		Email:    email,
	}

	if _, err := u.IsValid(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u User) IsValid() (bool, error) {
	if !isEmailValid(u.Email) {
		return false, ErrInvalidEmail
	}

	if _, err := uuid.Parse(u.ID); err != nil {
		return false, ErrInvalidUserID
	}

	if isEmpty(u.Username) {
		return false, ErrInvalidUsername
	}

	return true, nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func isEmpty(s string) bool {
	return len(s) == 0
}
