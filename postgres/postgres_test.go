package postgres

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestGetPostsByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "title", "content", "posted"}).
		AddRow(3, "title 1", "Content 1", now)

	query := "SELECT \\* FROM posts WHERE id=\\$1"
	prep := mock.ExpectPrepare(query)

	prep.ExpectQuery().WillReturnRows(rows)

	repository := NewPostgresRepository(db) // passes the mock to our code

	expectedPost := Post {
		ID: 1,
		Title: "title 1",
		Content: "Content 1",
		Posted: now,
	}
	
	actualPost, err := repository.GetByID(3)
	assert.NoError(t, err)
	assert.NotEmpty(t, actualPost)
	assert.ObjectsAreEqual(expectedPost, actualPost)
}