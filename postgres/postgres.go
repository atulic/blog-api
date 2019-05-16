package postgres

import (
	"database/sql"
	"fmt"
	"time"

	// postgres driver
	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

// New makes a new database using the connection string and
// returns it, otherwise returns the error
func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// Check that our connection is good
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

// BuildDbConnString builds the connection string to the db given a set of params
func BuildDbConnString(host string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password = %s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)
}

// Post shape
type Post struct {
	ID      int
	Title   string
	Content string
	Posted  time.Time
}

// GetPostByID is called within our post query for graphql
func (d *Db) GetPostByID(id int) []Post {
	// Prepare query, takes a name argument, protects from sql injection
	stmt, err := d.Prepare("SELECT * FROM posts WHERE id=$1")
	if err != nil {
		fmt.Println("GetPostByID Preparation Err: ", err)
	}

	// Make query with our stmt, passing in name argument
	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Println("GetPostByID Query Err: ", err)
	}

	// Create Post struct for holding each row's data
	var r Post
	// Create slice of Posts for our response
	posts := []Post{}
	// Copy the columns from row into the values pointed at by r (User)
	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Title,
			&r.Content,
			&r.Posted,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		posts = append(posts, r)
	}

	return posts
}

// CreatePost is called within our creation mutation for graphql
func (d *Db) CreatePost(post Post) {
	// Prepare query, takes a name argument, protects from sql injection
	stmt, err := d.Prepare("INSERT INTO posts (title, content, posted) VALUES ($1, $2, $3)")
	if err != nil {
		fmt.Println("CreatePost Preparation Err: ", err)
	}

	// Make query with our stmt, passing in name argument
	_, err = stmt.Exec(post.Title, post.Content, post.Posted)
	if err != nil {
		fmt.Println("GetPostByID Exec Err: ", err)
	}

	return
}
