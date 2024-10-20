package models

import "time"

type Product struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Qty       int       `db:"quantity"`
	Price     float64   `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
