package model

import "time"

type Post struct { 
	PID           string    `json:"-" db:"pid,omitempty"` 
	Title         string    `json:"title" db:"title,omitempty"`
	Summary       string    `json:"summary" db:"summary,omitempty"`
	Author        string    `json:"author" db:"author,omitempty"`
	Content       string    `json:"content" db:"content,omitempty"`
	CoverImage    string    `json:"cover_image" db:"cover_image,omitempty"`
	ContentImage  string    `json:"content_image" db:"content_image,omitempty"`
	CreatedAt     time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"-" db:"updated_at,omitempty"`
}
