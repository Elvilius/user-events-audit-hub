package grpc

import (
	"fmt"
	"net"

	"log/slog"

	eventgrpc "github.com/Elvilius/user-events-audit-hub/internal/grpc/event"
	service "github.com/Elvilius/user-events-audit-hub/internal/service/event"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(service *service.Service, port int) *App {
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	eventgrpc.Register(grpcServer, service)

	return &App{gRPCServer: grpcServer, port: port}
}

func (a *App) Run() {
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	slog.Info("grpc server started", slog.String("addr", list.Addr().String()))

	err = a.gRPCServer.Serve(list)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

}
