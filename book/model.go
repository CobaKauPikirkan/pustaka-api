package book

import "time"

type Book struct {
	ID          int			`json:"id"`
	Title       string		`json:"title"`
	Description string		`json:"description"`
	Price       int			`json:"price"`
	Discount 	int			`json:"discount"`
	Rating      int			`json:"rating"`
	CreatedAt   time.Time	`json:"created_at"`
	UpdatedAt   time.Time	`json:"updated_at"`
}