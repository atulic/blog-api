package postgres

// Repository represents the post repository contract
type Repository interface { 
	GetByID(id int) ([]Post)
	Update(a Post)
	Create(a Post)
	Delete(id int)
}