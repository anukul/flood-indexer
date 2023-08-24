package server

import (
	"github.com/anukul/flood-indexer/internal/config"
	"github.com/anukul/flood-indexer/internal/database"
)

type Server struct {
	cfg *config.Config
	db  *database.Client
}

func NewServer(cfg *config.Config, db *database.Client) *Server {
	return &Server{cfg: cfg, db: db}
}
