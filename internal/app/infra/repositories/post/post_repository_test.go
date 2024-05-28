package post

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/marcelocquadros/blog/internal/app/domain/post"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewPostRepository(sqlxDB)

	p := post.Post{
		ID:       "1",
		Title:    "Test Title",
		Content:  "Test Content",
		ImageURL: "http://example.com/image.jpg",
		UserID:   "user1",
	}

	mock.ExpectExec("INSERT INTO posts (id, title, content, image_url, user_id) VALUES(?, ?, ?, ?, ?)").
		WithArgs(p.ID, p.Title, p.Content, p.ImageURL, p.UserID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreatePost(p)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdatePost(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewPostRepository(sqlxDB)

	p := post.Post{
		ID:       "1",
		Title:    "Updated Title",
		Content:  "Updated Content",
		ImageURL: "http://example.com/updated_image.jpg",
	}

	mock.ExpectExec("UPDATE posts SET title=?, content=?, image_url=? WHERE id=?").
		WithArgs(p.Title, p.Content, p.ImageURL, p.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdatePost(p)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeletePost(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewPostRepository(sqlxDB)

	id := "1"

	mock.ExpectExec("DELETE FROM posts WHERE id=?").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeletePost(id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindAllPosts(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewPostRepository(sqlxDB)

	rows := sqlmock.NewRows([]string{"id", "title", "content", "image_url", "user_id"}).
		AddRow("1", "Title 1", "Content 1", "http://example.com/image1.jpg", "user1").
		AddRow("2", "Title 2", "Content 2", "http://example.com/image2.jpg", "user2")

	mock.ExpectQuery("SELECT * FROM posts").WillReturnRows(rows)

	posts, err := repo.FindAllPosts()
	assert.NoError(t, err)
	assert.Len(t, posts, 2)
	assert.Equal(t, "Title 1", posts[0].Title)
	assert.Equal(t, "Title 2", posts[1].Title)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindPostsByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewPostRepository(sqlxDB)

	id := "1"
	rows := sqlmock.NewRows([]string{"id", "title", "content", "image_url", "user_id"}).
		AddRow(id, "Title 1", "Content 1", "http://example.com/image1.jpg", "user1")

	mock.ExpectQuery("SELECT * FROM posts WHERE id=?").
		WithArgs(id).
		WillReturnRows(rows)

	p, err := repo.FindPostsByID(id)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, "Title 1", p.Title)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindPostsByIDNotFound(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewPostRepository(sqlxDB)

	id := "1"
	mock.ExpectQuery("SELECT * FROM posts WHERE id=?").
		WithArgs(id).
		WillReturnError(sql.ErrNoRows)

	p, err := repo.FindPostsByID(id)
	assert.NoError(t, err)
	assert.Nil(t, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}
