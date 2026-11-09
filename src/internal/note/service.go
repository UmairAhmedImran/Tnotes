package note

import (
	"context"
	"database/sql"
	"log/slog"
	"tnotes/internal/db/models"
	"tnotes/internal/users"
)

type Service struct {
	repo *Repository
	userSvc *users.Service
}

func NewService(db *sql.DB, userSvc *users.Service) *Service {
	return &Service{repo: NewRepository(db), userSvc: userSvc}
}


func (s *Service) AddNote(ctx context.Context, createInput CreateNote) error {

	user, err := s.userSvc.GetUser(ctx, models.User)

	if err != nil {
		return err
	}

	s.repo.Insert(ctx, createInput, user.ID)

	slog.Info("[NOTE][REPOSITORY] Note is added successfully")

	return nil
}
