package service

import (
	"context"

	"github.com/Elvilius/user-events-audit-hub/internal/domain/models"
	repo "github.com/Elvilius/user-events-audit-hub/internal/repo/event"
)

func NewService(repo *repo.Repo) *Service {
	return &Service{repo: repo}
}

type Service struct {
	repo *repo.Repo
}

func (s *Service) Create(ctx context.Context, event models.Event) (repo.EventID, error) {
	id, err := s.repo.CreateEvent(ctx, event)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *Service) List(ctx context.Context) ([]models.Event, error) {
	events, err := s.repo.GetEventList(ctx)
	if err != nil {
		return events, err
	}
	return events, nil
}
