package users

import "time"

type Signupreq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   string `db:"mobile"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Loginreq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Mobile    string    `db:"mobile"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
