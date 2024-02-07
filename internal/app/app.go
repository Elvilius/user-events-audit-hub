package app

import (
	"context"

	grpcapp "github.com/Elvilius/user-events-audit-hub/internal/app/grpc"
	"github.com/Elvilius/user-events-audit-hub/internal/config"
	"github.com/Elvilius/user-events-audit-hub/internal/lib/mongo"
	repo "github.com/Elvilius/user-events-audit-hub/internal/repo/event"
	service "github.com/Elvilius/user-events-audit-hub/internal/service/event"
)


type App struct {
	GRPCServer *grpcapp.App

}

func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	client, err := mongo.New(ctx, cfg.MongoUrl)  
	if err != nil {
		return nil, err
	}
	repo := repo.NewRepo(client)
	service := service.NewService(&repo)

	return &App{
		GRPCServer: grpcapp.New(service, cfg.GrpcPort),
	}, nil
}


