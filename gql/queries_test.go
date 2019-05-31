package gql

import (
	"testing"
	"time"

	mock "github.com/atulic/blog-api/mocks"
	"github.com/atulic/blog-api/postgres"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/graphql-go/graphql"
)

func TestFetchPost(t *testing.T) {
	fetchPostQuery := `{
		posts(id: 1) {
		  id
		  title
		  content
		  posted
		}
	  }`

	mockPost := postgres.Post{
		ID:      1,
		Title:   "Expected Title",
		Content: "Expected Content",
		Posted:  time.Now(),
	}

	// Create a new mock controller, which manages
	// our mock objects and their expectations
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockRepository(ctrl)

	// Ensure our mock repository is called once with the
	// expected params and set the mocks return value
	mockRepository.EXPECT().GetByID(1).Return(mockPost, nil).Times(1)

	rootQuery := NewRoot(mockRepository)

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query, Mutation: rootQuery.Mutation},
	)

	r := ExecuteQuery(fetchPostQuery, sc)

	expected := map[string]interface{}{
		"posts": map[string]interface{}{
			"id":      mockPost.ID,
			"content": mockPost.Content,
			"posted":  mockPost.Posted.String(),
			"title":   mockPost.Title}}

	assert.NoError(t, err)
	assert.Empty(t, r.Errors)
	assert.Equal(t, expected, r.Data)
}

func TestCreatePost(t *testing.T) {
	createPostMutation := `mutation {
		create(title: "Expected Title", content: "Expected Content") {
		  title
		}
	  }`

	expectedPost := postgres.Post{
		ID:      0,
		Title:   "Expected Title",
		Content: "Expected Content",
		Posted:  time.Now(),
	}

	// Create a new mock controller, which manages
	// our mock objects and their expectations
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockRepository(ctrl)

	// Ensure our mock repository is called once with the
	// expected params and set the mocks return value
	mockRepository.EXPECT().Create(expectedPost).Times(1)

	rootQuery := NewRoot(mockRepository)

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query, Mutation: rootQuery.Mutation},
	)

	r := ExecuteQuery(createPostMutation, sc)

	expected := map[string]interface{}{
		"posts": map[string]interface{}{
			"id":      expectedPost.ID,
			"content": expectedPost.Content,
			"posted":  expectedPost.Posted.String(),
			"title":   expectedPost.Title}}

	assert.NoError(t, err)
	assert.Empty(t, r.Errors)
	assert.Equal(t, expected, r.Data)
}
