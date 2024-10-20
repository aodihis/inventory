package models

import "time"

type User struct {
	ID        int        `db:"id"`
	Username  string     `db:"username"`
	Email     string     `db:"email"`
	Password  string     `db:"-"`
	CreatedAt time.Time  `db:"created_at"`
	LastLogin *time.Time `db:"last_login"`
}
