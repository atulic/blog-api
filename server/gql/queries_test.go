package gql

import (
	"testing"
	"time"

	mock "github.com/atulic/blog-api/mocks"
	"github.com/atulic/blog-api/postgres"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
)

func init() {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2014, time.September, 19, 16, 23, 53, 322, time.Local)
	})
}

func TestFetchPost(t *testing.T) {
	fetchPostQuery := `{
		posts(id: 1) {
		  id
		  title
		  content
		  posted
		}
	  }`

	mockPosts := createMockPostSlice()

	// Create a new mock controller, which manages
	// our mock objects and their expectations
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockRepository(ctrl)

	// Ensure our mock repository is called once with the
	// expected params and set the mocks return value
	mockRepository.EXPECT().GetByID(1).Return(mockPosts, nil).Times(1)

	rootQuery := NewRoot(mockRepository)

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query, Mutation: rootQuery.Mutation},
	)

	r := ExecuteQuery(fetchPostQuery, sc)

	newExpected := map[string]interface{}{
		"posts": []interface{}{map[string]interface{}{
			"content": mockPosts[0].Content,
			"id":      mockPosts[0].ID,
			"posted":  mockPosts[0].Posted.String(),
			"title":   mockPosts[0].Title}}}

	assert.NoError(t, err)
	assert.False(t, r.HasErrors())
	assert.Equal(t, newExpected, r.Data)
}

func TestGetAllPosts(t *testing.T) {
	fetchPostQuery := `{
		posts {
		  id
		  title
		  content
		  posted
		}
	  }`

	mockPosts := createMockPostSlice()

	// Create a new mock controller, which manages
	// our mock objects and their expectations
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockRepository(ctrl)

	// Ensure our mock repository is called once with the
	// expected params and set the mocks return value
	mockRepository.EXPECT().GetAllPosts().Return(mockPosts, nil).Times(1)

	rootQuery := NewRoot(mockRepository)

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query, Mutation: rootQuery.Mutation},
	)

	r := ExecuteQuery(fetchPostQuery, sc)

	newExpected := map[string]interface{}{
		"posts": []interface{}{map[string]interface{}{
			"content": mockPosts[0].Content,
			"id":      mockPosts[0].ID,
			"posted":  mockPosts[0].Posted.String(),
			"title":   mockPosts[0].Title}}}

	assert.NoError(t, err)
	assert.False(t, r.HasErrors())
	assert.Equal(t, newExpected, r.Data)
}

func TestCreatePost(t *testing.T) {
	createPostMutation := `mutation {
		create(title: "Expected Title", content: "Expected Content") {
		  id
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

	expected := map[string]interface{}(
		map[string]interface{}{
			"create": map[string]interface{}{
				"id":    0,
				"title": "Expected Title"}})

	assert.NoError(t, err)
	assert.False(t, r.HasErrors())
	assert.Equal(t, expected, r.Data)
}

func TestDeletePost(t *testing.T) {
	deletePostMutation := `mutation {
		delete(id:1) {
		  title
		}
	  }`

	// Create a new mock controller, which manages
	// our mock objects and their expectations
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockRepository(ctrl)

	// Ensure our mock repository is called once with the
	// expected params and set the mocks return value
	mockRepository.EXPECT().Delete(1).Times(1)

	rootQuery := NewRoot(mockRepository)

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query, Mutation: rootQuery.Mutation},
	)

	r := ExecuteQuery(deletePostMutation, sc)

	assert.NoError(t, err)
	assert.False(t, r.HasErrors())
}

func TestUpdatePost(t *testing.T) {
	updatePostMutation := `mutation {
		update(id: 1, title: "Updated title", content: "Updated content :)") {
		  title
		  content
		}
	  }`

	updatedPost := postgres.Post{
		ID:      1,
		Title:   "Updated title",
		Content: "Updated content :)",
		Posted:  time.Now(),
	}

	// Create a new mock controller, which manages
	// our mock objects and their expectations
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockRepository(ctrl)

	// Ensure our mock repository is called once with the
	// expected params and set the mocks return value
	mockRepository.EXPECT().Update(updatedPost).Times(1)

	rootQuery := NewRoot(mockRepository)

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query, Mutation: rootQuery.Mutation},
	)

	r := ExecuteQuery(updatePostMutation, sc)

	expectedResult := map[string]interface{}(
		map[string]interface{}{
			"update": map[string]interface{}{
				"title":   "Updated title",
				"content": "Updated content :)"}})

	assert.NoError(t, err)
	assert.False(t, r.HasErrors())
	assert.Equal(t, expectedResult, r.Data)
}

func createMockPostSlice() []postgres.Post {
	return []postgres.Post{{
		ID:      1,
		Title:   "Expected Title",
		Content: "Expected Content",
		Posted:  time.Now(),
	},
	}
}
