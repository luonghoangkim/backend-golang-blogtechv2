package model

import "time"

type User struct {
	UserID    string `json:"-" db:"user_id, omitempty"`
	FullName  string `json:"fullName,omitempty" db:"full_name, omitempty"`
	Email     string `json:"email,omitempty" db:"email, omitempty"`
	PassWord  string `json:"-" db:"password, omitempty"`
	Role      string `json:"-" db:"role, omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at, omitempty"`
	Token string  `json:"token,omitempty"`
}

func New() User {
	return User{}
}

func (u User) GetEmail() string {
	return u.Email
}