package app

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcHandler "github.com/Elvilius/user-events-audit-hub/internal/handlers/grpc"
	httpHandler "github.com/Elvilius/user-events-audit-hub/internal/handlers/http"

	"github.com/Elvilius/user-events-audit-hub/internal/config"
	"github.com/Elvilius/user-events-audit-hub/internal/lib/mongo"
	repo "github.com/Elvilius/user-events-audit-hub/internal/repo/event"
	service "github.com/Elvilius/user-events-audit-hub/internal/service/event"
)

type App struct {
	GRPCServer *grpc.Server
	HTTPServer *http.Server
}

func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	client, err := mongo.New(ctx, cfg.MongoDns)
	if err != nil {
		return nil, err
	}
	repo := repo.NewRepo(client)
	service := service.NewService(repo)

	return &App{
		GRPCServer: newGrpc(cfg, service),
		HTTPServer: newHttp(cfg, service),
	}, nil
}

func newGrpc(cfg *config.Config, service *service.Service) *grpc.Server {
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	grpcHandler.Register(grpcServer, service)
	return grpcServer
}

func newHttp(cfg *config.Config, service *service.Service) *http.Server {
	mux := http.NewServeMux()
	httpHandler.Register(mux, service)

	httpServer := &http.Server{
		Handler: mux,
		Addr:    ":" + fmt.Sprint(cfg.HttpPort),
	}
	return httpServer
}
