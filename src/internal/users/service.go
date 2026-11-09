package users

import (
	"context"
	"database/sql"
	"tnotes/internal/db/models"
)

type Service struct {
	repo *Repository
}

func NewService (db *sql.DB) *Service {
	return &Service{repo: NewRepository(db)}
}

func (s *Service) GetUser(ctx context.Context, username models.Username) (models.Users, error) {
	return s.repo.RetrieveUser(ctx, username);
}