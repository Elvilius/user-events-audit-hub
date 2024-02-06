package app

import (
	"context"
	"log"
	"net"

	repo "github.com/Elvilius/user-events-audit-hub/internal/repo/event"
	service "github.com/Elvilius/user-events-audit-hub/internal/service/event"
	desc "github.com/Elvilius/user-events-audit-hub/proto/event_v1"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)


type App struct {
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	client, err := newClient("mongodb://event:secret@eventsdb:27017/events?authSource=admin&directConnection=true", ctx)
	if err != nil {
		return nil, err
	}
	repo := repo.NewRepo(client)
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcServer)
	desc.RegisterEventV1Server(a.grpcServer, &EventServerApi{EventService: service.NewService(&repo)})
	return a, nil
}

func (a *App) Run() error {
	return a.runGRPCServer()
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", ":9000")

	list, err := net.Listen("tcp", ":9000")
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func newClient(url string, ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
