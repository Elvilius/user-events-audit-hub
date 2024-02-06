package service

import (
	"context"

	repo "github.com/Elvilius/user-events-audit-hub/internal/repo/event"
)

type CreateEventDto struct {
	UserId     int
	EventType  string
	SystemName string
	Message     string
	Severity    string
	Metadata    map[string]string
}



type EventIdDto struct {
	Id string
}

type Service struct {
	repo *repo.Repo
}

func (s *Service) Create(ctx context.Context, createDto CreateEventDto) (EventIdDto, error) {
	id, err := s.repo.CreateEvent(ctx, repo.CreateEventDto(createDto))
	if err != nil {
		return EventIdDto{}, err
	}
	return EventIdDto{Id: id}, nil
}

func NewService(repo *repo.Repo) Service {
	return Service{repo: repo}
}
