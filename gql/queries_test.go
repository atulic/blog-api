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

var fetchPostQuery = `{
	posts(id: 1) {
	  id
	  title
	  content
	  posted
	}
  }`

func TestFetchPost(t *testing.T) {
	now := time.Now()

	mockPost := postgres.Post{
		ID:      1,
		Title:   "Expected Title",
		Content: "Expected Content",
		Posted:  now,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockRepository(ctrl)
	mockRepository.EXPECT().GetByID(1).Return(mockPost, nil).Times(1)

	rootQuery := NewRoot(mockRepository)

	// Create a new graphql schema, passing in the the root query
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
