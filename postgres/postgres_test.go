package postgres

import (
	"testing"
	"time"
	"database/sql"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setupMocks(t *testing.T) (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock, err
}


func TestBuildsConnectionString(t *testing.T){
	dbStringExpected := "host=localhost port=4000 user=user password = password dbname=db_name sslmode=disable"
	dbStringActual := BuildDbConnString("localhost", 4000, "user", "password", "db_name")

	assert.Equal(t, dbStringExpected, dbStringActual, "The connection string matches the expected")
}

func TestGetPostsByID(t *testing.T) {

	db, mock, err := setupMocks(t)
	defer db.Close()

	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "title", "content", "posted"}).
		AddRow(3, "title 1", "Content 1", now)

	query := "SELECT \\* FROM posts WHERE id=\\$1"
	prep := mock.ExpectPrepare(query)

	prep.ExpectQuery().WillReturnRows(rows)

	repository := NewPostgresRepository(db) // passes the mock to our code

	expectedPost := Post{
		ID:      1,
		Title:   "title 1",
		Content: "Content 1",
		Posted:  now,
	}

	actualPost, err := repository.GetByID(3)
	assert.NoError(t, err)
	assert.NotEmpty(t, actualPost)
	assert.ObjectsAreEqual(expectedPost, actualPost)
}

func TestCreatePost(t *testing.T) {
	db, mock, err := setupMocks(t)
	defer db.Close()

	now := time.Now()

	post := Post{
		ID:      1,
		Title:   "title 1",
		Content: "Content 1",
		Posted:  now,
	}

	query := "INSERT INTO posts"
	mock.ExpectPrepare(query).ExpectExec().WithArgs(post.Title, post.Content, post.Posted).WillReturnResult(sqlmock.NewResult(1, 1))

	repository := NewPostgresRepository(db) // passes the mock to our code

	err = repository.Create(post)
	assert.NoError(t, err)
}

func TestUpdatePost(t *testing.T) {
	db, mock, err := setupMocks(t)
	defer db.Close()

	now := time.Now()

	post := Post{
		ID:      1,
		Title:   "title 1",
		Content: "Content 1",
		Posted:  now,
	}

	query := "UPDATE posts"
	mock.ExpectPrepare(query).ExpectExec().WithArgs(post.ID, post.Title, post.Content, post.Posted).WillReturnResult(sqlmock.NewResult(1, 1))

	repository := NewPostgresRepository(db) // passes the mock to our code

	err = repository.Update(post)
	assert.NoError(t, err)
}

func TestDeletePost(t *testing.T) {
	db, mock, err := setupMocks(t)
	defer db.Close()

	now := time.Now()

	post := Post{
		ID:      1,
		Title:   "title 1",
		Content: "Content 1",
		Posted:  now,
	}

	query := "DELETE FROM posts"
	mock.ExpectPrepare(query).ExpectExec().WithArgs(post.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	repository := NewPostgresRepository(db) // passes the mock to our code

	err = repository.Delete(post.ID)
	assert.Empty(t, err)
}
