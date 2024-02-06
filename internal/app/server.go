package app

import (
	"context"

	service "github.com/Elvilius/user-events-audit-hub/internal/service/event"
	e "github.com/Elvilius/user-events-audit-hub/proto/event_v1"
)

type EventServerApi struct {
	e.UnimplementedEventV1Server
	EventService service.Service
}

func (s *EventServerApi) Create(ctx context.Context, req *e.CreateRequest) (*e.CreateResponse, error) {
	//TODO Добавить валидацию
	// Не понятно как передавать данные от grpc до с
	id, err := s.EventService.Create(ctx, service.CreateEventDto{UserId: int(req.Event.UserId)})

	if err != nil {
		return &e.CreateResponse{}, err
	}
	return &e.CreateResponse{Id: id.Id}, nil
}
