package models

import "time"

type Person struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Gift struct {
	ID          string    `json:"id" db:"id"`
	PersonID    string    `json:"person_id" db:"person_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	IsSelected  bool      `json:"is_selected" db:"is_selected"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type CreatePersonInput struct {
	Name string `json:"name" binding:"required"`
}

type CreateGiftInput struct {
	PersonID    string `json:"person_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateGiftInput struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	IsSelected  *bool   `json:"is_selected,omitempty"`
}
