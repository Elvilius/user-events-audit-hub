package event

import (
	"encoding/json"
	"net/http"

	service "github.com/Elvilius/user-events-audit-hub/internal/service/event"
)

type EventServerApi struct {
	EventService *service.Service
}

func Register(server *http.ServeMux, service *service.Service) {
	s := EventServerApi{EventService: service}
	server.HandleFunc("/api/event/list", s.List)
}

func (e *EventServerApi) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	events, err := e.EventService.List(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
