package user

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/marcelocquadros/blog/internal/app/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewUserRepository(sqlxDB)

	u := &user.User{
		ID:       "1",
		Username: "testuser",
		Email:    "testuser@example.com",
	}

	mock.ExpectExec("INSERT INTO users(id, username, email) VALUES (?, ?, ?)").
		WithArgs(u.ID, u.Username, u.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateUser(u)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

	// Error case
	mock.ExpectExec("INSERT INTO users(id, username, email) VALUES (?, ?, ?)").
		WithArgs(u.ID, u.Username, u.Email).
		WillReturnError(errors.New("insert error"))

	err = repo.CreateUser(u)
	assert.Error(t, err)
	assert.Equal(t, "insert error", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewUserRepository(sqlxDB)

	id := "1"

	mock.ExpectExec("DELETE FROM users WHERE id=?").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteUser(id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

	// Error case
	mock.ExpectExec("DELETE FROM users WHERE id=?").
		WithArgs(id).
		WillReturnError(errors.New("delete error"))

	err = repo.DeleteUser(id)
	assert.Error(t, err)
	assert.Equal(t, "delete error", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewUserRepository(sqlxDB)

	u := &user.User{
		ID:       "1",
		Username: "updateduser",
		Email:    "updateduser@example.com",
	}

	mock.ExpectExec("UPDATE users SET username=?, email=?").
		WithArgs(u.Username, u.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateUser(u)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

	// Error case
	mock.ExpectExec("UPDATE users SET username=?, email=?").
		WithArgs(u.Username, u.Email).
		WillReturnError(errors.New("update error"))

	err = repo.UpdateUser(u)
	assert.Error(t, err)
	assert.Equal(t, "update error", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewUserRepository(sqlxDB)

	rows := sqlmock.NewRows([]string{"id", "username", "email"}).
		AddRow("1", "user1", "user1@example.com").
		AddRow("2", "user2", "user2@example.com")

	mock.ExpectQuery("SELECT * FROM users").WillReturnRows(rows)

	users, err := repo.FindAllUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "user1", users[0].Username)
	assert.Equal(t, "user2", users[1].Username)
	assert.NoError(t, mock.ExpectationsWereMet())

	// Error case
	mock.ExpectQuery("SELECT * FROM users").
		WillReturnError(errors.New("select error"))

	users, err = repo.FindAllUsers()
	assert.Error(t, err)
	assert.Equal(t, "select error", err.Error())
	assert.Empty(t, users)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindUserByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewUserRepository(sqlxDB)

	id := "1"
	rows := sqlmock.NewRows([]string{"id", "username", "email"}).
		AddRow(id, "user1", "user1@example.com")

	mock.ExpectQuery("SELECT * FROM users WHERE id=?").
		WithArgs(id).
		WillReturnRows(rows)

	u, err := repo.FindUserByID(id)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, "user1", u.Username)
	assert.NoError(t, mock.ExpectationsWereMet())

	// Error case
	mock.ExpectQuery("SELECT * FROM users WHERE id=?").
		WithArgs(id).
		WillReturnError(errors.New("select error"))

	u, err = repo.FindUserByID(id)
	assert.Error(t, err)
	assert.Nil(t, u)
	assert.Equal(t, "select error", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindUserByIDNotFound(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewUserRepository(sqlxDB)

	id := "1"
	mock.ExpectQuery("SELECT * FROM users WHERE id=?").
		WithArgs(id).
		WillReturnError(sql.ErrNoRows)

	u, err := repo.FindUserByID(id)
	assert.NoError(t, err)
	assert.Nil(t, u)
	assert.NoError(t, mock.ExpectationsWereMet())
}
