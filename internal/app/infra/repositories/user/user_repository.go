package user

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/marcelocquadros/blog/internal/app/domain/user"
)

type (
	userRepository struct {
		db *sqlx.DB
	}
)

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur userRepository) CreateUser(u *user.User) error {
	if _, err := ur.db.Exec("INSERT INTO users(id, username, email) VALUES (?, ?, ?)", u.ID, u.Username, u.Email); err != nil {
		return err
	}

	return nil
}

func (ur userRepository) DeleteUser(id string) error {
	if _, err := ur.db.Exec("DELETE FROM users WHERE id=?", id); err != nil {
		return err
	}

	return nil
}

func (ur userRepository) UpdateUser(u *user.User) error {
	if _, err := ur.db.Exec("UPDATE users SET username=?, email=?", u.Username, u.Email); err != nil {
		return err
	}
	return nil
}

func (ur userRepository) FindAllUsers() ([]user.User, error) {
	users := []user.User{}
	if err := ur.db.Select(&users, "SELECT * FROM users"); err != nil {
		return []user.User{}, err
	}

	return users, nil
}

func (ur userRepository) FindUserByID(id string) (*user.User, error) {
	user := user.User{}
	if err := ur.db.Get(&user, "SELECT * FROM users WHERE id=?", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
