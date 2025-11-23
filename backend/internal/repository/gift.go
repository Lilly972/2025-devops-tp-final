package repository

import (
	"context"
	"fmt"

	"github.com/anthonyquere/christmas-gift-list/internal/database"
	"github.com/anthonyquere/christmas-gift-list/internal/models"
	"github.com/google/uuid"
)

type GiftRepository struct{}

func NewGiftRepository() *GiftRepository {
	return &GiftRepository{}
}

func (r *GiftRepository) GetByPersonID(ctx context.Context, personID string) ([]models.Gift, error) {
	query := `SELECT id, person_id, title, description, is_selected, created_at 
	          FROM gifts WHERE person_id = $1 ORDER BY created_at DESC`

	rows, err := database.DB.Query(ctx, query, personID)
	if err != nil {
		return nil, fmt.Errorf("failed to query gifts: %w", err)
	}
	defer rows.Close()

	gifts := make([]models.Gift, 0)
	for rows.Next() {
		var g models.Gift
		if err := rows.Scan(&g.ID, &g.PersonID, &g.Title, &g.Description, &g.IsSelected, &g.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan gift: %w", err)
		}
		gifts = append(gifts, g)
	}

	return gifts, nil
}

func (r *GiftRepository) Create(ctx context.Context, input models.CreateGiftInput) (*models.Gift, error) {
	id := uuid.New().String()
	query := `INSERT INTO gifts (id, person_id, title, description, is_selected) 
	          VALUES ($1, $2, $3, $4, false) 
	          RETURNING id, person_id, title, description, is_selected, created_at`

	var g models.Gift
	err := database.DB.QueryRow(ctx, query, id, input.PersonID, input.Title, input.Description).
		Scan(&g.ID, &g.PersonID, &g.Title, &g.Description, &g.IsSelected, &g.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create gift: %w", err)
	}

	return &g, nil
}

func (r *GiftRepository) Update(ctx context.Context, id string, input models.UpdateGiftInput) (*models.Gift, error) {
	query := `UPDATE gifts SET 
	          title = COALESCE($2, title),
	          description = COALESCE($3, description),
	          is_selected = COALESCE($4, is_selected)
	          WHERE id = $1
	          RETURNING id, person_id, title, description, is_selected, created_at`

	var g models.Gift
	err := database.DB.QueryRow(ctx, query, id, input.Title, input.Description, input.IsSelected).
		Scan(&g.ID, &g.PersonID, &g.Title, &g.Description, &g.IsSelected, &g.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update gift: %w", err)
	}

	return &g, nil
}

func (r *GiftRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM gifts WHERE id = $1`

	_, err := database.DB.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete gift: %w", err)
	}

	return nil
}

func (r *GiftRepository) SelectGift(ctx context.Context, personID, giftID string) (*models.Gift, error) {
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Unselect all gifts for this person
	_, err = tx.Exec(ctx, `UPDATE gifts SET is_selected = false WHERE person_id = $1`, personID)
	if err != nil {
		return nil, fmt.Errorf("failed to unselect gifts: %w", err)
	}

	// Select the specified gift
	query := `UPDATE gifts SET is_selected = true WHERE id = $1 AND person_id = $2 
	          RETURNING id, person_id, title, description, is_selected, created_at`

	var g models.Gift
	err = tx.QueryRow(ctx, query, giftID, personID).
		Scan(&g.ID, &g.PersonID, &g.Title, &g.Description, &g.IsSelected, &g.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to select gift: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &g, nil
}
