package main

import (
	"context"
	"github.com/Pasca11/GoBlogApp/internal/api/server"
	"github.com/Pasca11/GoBlogApp/internal/config"
	"log"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	srv := server.New(server.WithConfig(&cfg.Server))

	log.Fatal(srv.Run(ctx))
}
