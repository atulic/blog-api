package postgres

// Repository represents the post repository contract
type Repository interface { 
	GetByID(id int) ([]Post, error)
	Update(a Post) error
	Create(a Post) error
	Delete(id int) error
}