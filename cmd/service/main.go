package main

import (
	"context"
	"log/slog"

	"github.com/Elvilius/user-events-audit-hub/internal/config"
	// "log"
	"github.com/Elvilius/user-events-audit-hub/internal/app"
)

func main() {
	ctx := context.Background()
	cfg := config.Load()

	done := make(chan struct{})

	a, err := app.NewApp(ctx, cfg)

	if err != nil {
		slog.Error(err.Error())
	}

	go func() {
		a.HTTPServer.Run()
	}()

	go func() {
		a.GRPCServer.Run()
	}()
	<-done
}
