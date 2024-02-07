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
