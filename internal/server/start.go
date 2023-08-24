package server

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/leaderboard_stats", s.GetLeaderboardStats)
	handler := cors.AllowAll().Handler(mux)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.ApiPort), handler)
}
