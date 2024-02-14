package grpc

import (
	"net/http"

	"log/slog"

	eventhttp "github.com/Elvilius/user-events-audit-hub/internal/http/event"

	service "github.com/Elvilius/user-events-audit-hub/internal/service/event"
)

type App struct {
	httpServer   *http.Server
	http_address string
}

func New(service *service.Service, http_address string) *App {
	mux := http.NewServeMux()
	eventhttp.Register(mux, service)

	httpServer := &http.Server{
		Handler: mux,
		Addr:    http_address,
	}

	return &App{httpServer: httpServer, http_address: http_address}
}

func (a *App) Run() {
	slog.Info("Http server started", slog.String("addr", a.http_address))
	err := a.httpServer.ListenAndServe()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

}
