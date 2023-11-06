package entity

import "time"

type Product struct {
	ID         int
	Title      string
	Price      int
	Stock      int
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
