package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/anukul/flood-indexer/internal/config"
	"github.com/anukul/flood-indexer/internal/database"
	"github.com/anukul/flood-indexer/internal/indexer"
	"github.com/anukul/flood-indexer/internal/server"
	"golang.org/x/exp/slog"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	db, err := database.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	idx, err := indexer.NewIndexer(cfg, db)
	if err != nil {
		panic(err)
	}
	go func() {
		slog.Info("starting indexer")
		err := idx.Start(ctx)
		slog.Info("indexer stopped")
		if err != nil {
			panic(err)
		}
	}()

	srv := server.NewServer(cfg, db)
	go func() {
		slog.Info("starting server")
		err := srv.Start()
		slog.Info("server stopped")
		if err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()
}
