package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Elvilius/user-events-audit-hub/internal/app"
	"github.com/Elvilius/user-events-audit-hub/internal/config"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		slog.Error(err.Error())
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(
		sigCh, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM,
	)
	a, err := app.NewApp(ctx, cfg)
	if err != nil {
		slog.Error(err.Error())
	}

	go func() {
		slog.Info("Http server started", slog.String("port", fmt.Sprintf("%d", cfg.HttpPort)))
		err := a.HTTPServer.ListenAndServe()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}()

	go func() {
		list, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GrpcPort))
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		slog.Info("grpc server started", slog.String("addr", list.Addr().String()))
		err = a.GRPCServer.Serve(list)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}()

	defer func() {
		a.GRPCServer.GracefulStop()
		if err := a.HTTPServer.Shutdown(ctx); err != nil {
			slog.Error(err.Error())
		}
	}()
	<-sigCh
}
