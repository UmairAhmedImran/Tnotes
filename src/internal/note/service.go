package note

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	"tnotes/internal/db/models"
	"tnotes/internal/users"
)

type Service struct {
	repo    *Repository
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

func (s *Service) ListNote(ctx context.Context, detailed bool) error {
	notes, err := s.repo.List(ctx)

	if err != nil {
		return err
	}

	for i := range notes {
		log.Println()
		log.Println("-------------------------------------------")
		log.Printf("Title: %s", notes[i].Title)
		if detailed {
			log.Printf("Content: %s", notes[i].Content)
		}
		log.Printf("Created: %s", notes[i].CreatedAt)
		log.Println("-------------------------------------------")
		log.Println()
	}

	return nil
}
