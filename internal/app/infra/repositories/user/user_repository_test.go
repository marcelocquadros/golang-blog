package user

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/marcelocquadros/blog/internal/app/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var dbMock = sqlx.NewDb(db, "sqlmock")
	defer db.Close()

	if err != nil {
		t.Fail()
	}
	repo := NewUserRepository(dbMock)
	user := &user.User{ID: "c93f0ec6-70af-4598-b767-dadd1554be3d", Username: "marcelo", Email: "marcelo@gmail.com"}

	t.Run("Create user success", func(t *testing.T) {

		mock.ExpectExec("INSERT INTO users(id, username, email) VALUES (?, ?, ?)").
			WithArgs(user.ID, user.Username, user.Email).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.CreateUser(user)
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Create user fail", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO users(id, username, email) VALUES (?, ?, ?)").
			WithArgs(user.ID, user.Username, user.Email).
			WillReturnError(errors.New("error"))

		err = repo.CreateUser(user)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})

	t.Run("Delete user success", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM users WHERE id=?").WithArgs(user.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.DeleteUser(user.ID)
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Delete user fail", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM users WHERE id=?").
			WithArgs(user.ID).
			WillReturnError(errors.New("error"))

		err = repo.DeleteUser(user.ID)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Find user by id fail", func(t *testing.T) {

		mock.ExpectQuery("SELECT * FROM users WHERE id=?").
			WithArgs(user.ID).
			WillReturnError(errors.New("error"))

		_, err := repo.FindUserByID(user.ID)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Find user not found", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM users WHERE id=?").
			WithArgs(user.ID).
			WillReturnError(sql.ErrNoRows)

		_, err = repo.FindUserByID(user.ID)
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Find user by id success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email"}).
			AddRow(user.ID, user.Username, user.Email)

		mock.ExpectQuery("SELECT * FROM users WHERE id=?").WithArgs(user.ID).
			WillReturnRows(rows)

		user, _ := repo.FindUserByID(user.ID)

		assert.NotNil(t, user)
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Find all users success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email"}).
			AddRow(user.ID, user.Username, user.Email)

		mock.ExpectQuery("SELECT * FROM users").
			WithoutArgs().
			WillReturnRows(rows)

		users, err := repo.FindAllUsers()

		assert.NoError(t, err)
		assert.NotEmpty(t, users)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})

	t.Run("Find all users fail", func(t *testing.T) {

		mock.ExpectQuery("SELECT * FROM users").
			WithoutArgs().
			WillReturnError(errors.New("error"))

		users, err := repo.FindAllUsers()

		assert.Error(t, err)
		assert.Empty(t, users)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})

	t.Run("Update user success", func(t *testing.T) {

		mock.ExpectExec("UPDATE users SET username=?, email=?").
			WithArgs(user.Username, user.Email).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.UpdateUser(user)

		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})

	t.Run("Update user fail", func(t *testing.T) {

		mock.ExpectExec("UPDATE users SET username=?, email=?").
			WithArgs(user.Username, user.Email).
			WillReturnError(errors.New("error"))

		err := repo.UpdateUser(user)

		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})
}
