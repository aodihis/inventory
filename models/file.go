package models

import "time"

type File struct {
	ID        int       `db:"id"`
	ProductID int       `db:"product_id"`
	FileName  string    `db:"file_name"`
	FilePath  string    `db:"file_path"`
	Folder    *string   `db:"folder"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
