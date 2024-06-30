package users

import "time"

type Signupreq struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type Loginreq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        int       `db:"id"`
	FirstName string    `db:"firstname"`
	LastName  string    `db:"lastname"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Mobile    string    `db:"mobile"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
