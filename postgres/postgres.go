package postgres

import (
	"database/sql"
	"fmt"
	"time"

	// postgres driver
	_ "github.com/lib/pq"
)

// DbConnection represents a SQL connection to the db. 
// Used to interact with the real database
type DbConnection struct {
	*sql.DB
}

// NewPostgresRepository takes a database connection and
// returns our repository
func NewPostgresRepository(Conn *sql.DB) Repository {
	return &DbConnection{Conn}
}

// New makes a new database using the connection string and
// returns it, otherwise returns the error
func New(connString string) (*DbConnection, error) {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	// Check that our connection is good
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DbConnection{db}, nil
}

// BuildDbConnString builds the connection string to the db given a set of params
func BuildDbConnString(host string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password = %s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)
}

// Post model
type Post struct {
	ID      int
	Title   string
	Content string
	Posted  time.Time
}

// GetByID is called within our post query for graphql
func (d *DbConnection) GetByID(id int) ([]Post, error) {
	// Prepare query, takes a id argument, protects from sql injection
	stmt, err := d.Prepare("SELECT * FROM posts WHERE id=$1")
	if err != nil {
		fmt.Println("GetPostByID Preparation Err: ", err)
	}

	// Make query with our stmt, passing in id argument
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

	return posts, err
}

// Create a new record in the DB with an auto-incrementing ID
func (d *DbConnection) Create(post Post)(error) {
	// Prepare query, takes a name argument, protects from sql injection
	stmt, err := d.Prepare("INSERT INTO posts (title, content, posted) VALUES ($1, $2, $3)")
	if err != nil {
		fmt.Println("CreatePost Preparation Err: ", err)
	}

	// Make query with our stmt, passing in name argument
	_, err = stmt.Exec(post.Title, post.Content, post.Posted)
	if err != nil {
		fmt.Println("CreatePost Exec Err: ", err)
	}

	return err
}

// Update the post in the DB where the ID matches
func (d *DbConnection) Update(post Post)(error) {
	// Prepare query, takes a name argument, protects from sql injection
	stmt, err := d.Prepare("UPDATE posts SET title = $2, content = $3, posted = $4 WHERE id = $1")
	if err != nil {
		fmt.Println("UpdatePost Preparation Err: ", err)
	}

	// Make query with our stmt, passing in name argument
	_, err = stmt.Exec(post.ID, post.Title, post.Content, post.Posted)
	if err != nil {
		fmt.Println("UpdatePost Exec Err: ", err)
	}

	return err
}

// Delete a post from the DB with a given ID
func (d *DbConnection) Delete(id int) error {
	// Prepare query, takes a name argument, protects from sql injection
	stmt, err := d.Prepare("DELETE FROM posts WHERE id = $1")
	if err != nil {
		fmt.Println("UpdatePost Preparation Err: ", err)
	}

	// Make query with our stmt, passing in name argument
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Println("UpdatePost Exec Err: ", err)
	}

	return err
}
