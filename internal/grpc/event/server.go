package event

import (
	"context"
	"fmt"

	desc "github.com/Elvilius/user-events-audit-hub/api/grpc/proto/event_v1"

	e "github.com/Elvilius/user-events-audit-hub/api/grpc/proto/event_v1"
	"github.com/Elvilius/user-events-audit-hub/internal/domain/models"
	service "github.com/Elvilius/user-events-audit-hub/internal/service/event"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EventServerApi struct {
	e.UnimplementedEventV1Server
	EventService *service.Service
}

func Register(grpc *grpc.Server, service *service.Service) {
	desc.RegisterEventV1Server(grpc, &EventServerApi{EventService: service})
}

func (s *EventServerApi) Create(ctx context.Context, req *e.CreateRequest) (*e.CreateResponse, error) {
	errors := validateCreate(req)

	if len(errors)!= 0 {
		return nil, fmt.Errorf("%v", errors)
	}

	id, err := s.EventService.Create(ctx, models.Event{
		UserId:     int(req.Event.GetUserId()),
		Message:    req.Event.GetMessage(),
		SystemName: req.Event.GetSystemName(),
		Metadata:   req.Event.GetMetadata(),
		Severity:   req.Event.GetSeverity(),
		EventType:  req.Event.GetEventType(),
	})

	if err != nil {
		return &e.CreateResponse{}, err
	}
	return &e.CreateResponse{Id: string(id)}, nil
}

func validateCreate(req *e.CreateRequest) []error {
	 var errors []error
	
	if req.Event.GetUserId() == 0 {
		 errors =  append(errors, status.Error(codes.InvalidArgument, "userId is required"))
	}
	if req.Event.GetMessage() == "" {
		errors =  append(errors, status.Error(codes.InvalidArgument, "message is required"))
	}
	if req.Event.GetEventType() == "" {
		errors =  append(errors, status.Error(codes.InvalidArgument, "event_type is required"))
	}
	if req.Event.GetSeverity() == "" {
		errors =  append(errors, status.Error(codes.InvalidArgument, "severity is required"))
	}
	if req.Event.GetSystemName() == "" {
		errors =  append(errors, status.Error(codes.InvalidArgument, "system_name is required"))
	}

	return errors

}
