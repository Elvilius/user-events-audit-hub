package event

import (
	"encoding/json"
	"fmt"
	"net/http"

	service "github.com/Elvilius/user-events-audit-hub/internal/service/event"
)

type httpHandler struct {
	EventService *service.Service
}

func Register(server *http.ServeMux, service *service.Service) {
	h := httpHandler{EventService: service}
	server.HandleFunc("/api/event/list", h.List)
}

func (h *httpHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	events, err := h.EventService.List(ctx)
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
	n, err := w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(n)
}
