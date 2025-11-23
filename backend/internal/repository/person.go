package repository

import (
	"context"
	"fmt"

	"github.com/anthonyquere/christmas-gift-list/internal/database"
	"github.com/anthonyquere/christmas-gift-list/internal/models"
	"github.com/google/uuid"
)

type PersonRepository struct{}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{}
}

func (r *PersonRepository) GetAll(ctx context.Context) ([]models.Person, error) {
	query := `SELECT id, name, created_at FROM people ORDER BY created_at DESC`

	rows, err := database.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query people: %w", err)
	}
	defer rows.Close()

	people := make([]models.Person, 0)
	for rows.Next() {
		var p models.Person
		if err := rows.Scan(&p.ID, &p.Name, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan person: %w", err)
		}
		people = append(people, p)
	}

	return people, nil
}

func (r *PersonRepository) GetByID(ctx context.Context, id string) (*models.Person, error) {
	query := `SELECT id, name, created_at FROM people WHERE id = $1`

	var p models.Person
	err := database.DB.QueryRow(ctx, query, id).Scan(&p.ID, &p.Name, &p.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get person: %w", err)
	}

	return &p, nil
}

func (r *PersonRepository) Create(ctx context.Context, input models.CreatePersonInput) (*models.Person, error) {
	id := uuid.New().String()
	query := `INSERT INTO people (id, name) VALUES ($1, $2) RETURNING id, name, created_at`

	var p models.Person
	err := database.DB.QueryRow(ctx, query, id, input.Name).Scan(&p.ID, &p.Name, &p.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create person: %w", err)
	}

	return &p, nil
}

func (r *PersonRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM people WHERE id = $1`

	_, err := database.DB.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete person: %w", err)
	}

	return nil
}
