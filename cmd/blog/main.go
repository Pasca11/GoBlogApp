package main

import (
	"context"
	"github.com/Pasca11/GoBlogApp/internal/api/server"
	"github.com/Pasca11/GoBlogApp/internal/config"
	"github.com/Pasca11/GoBlogApp/internal/pkg/logger"
	"time"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	log, err := logger.New(cfg.Logger)
	if err != nil {
		panic(err)
	}

	srv := server.New(server.WithConfig(cfg.Server))
	log.InfoContext(ctx, "starting server")
	go func() {
		if err := srv.Run(ctx); err != nil {
			log.ErrorContext(ctx, "error starting server")
		}
	}()

	time.Sleep(10 * time.Minute)
}
