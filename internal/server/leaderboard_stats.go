package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/exp/slog"

	"github.com/anukul/flood-indexer/internal/types"
)

type LeaderboardStatsResponse struct {
	Stats []*types.UserStats `json:"stats"`
}

func getRangeStartTime(rawRangeSeconds string) (time.Time, error) {
	if rawRangeSeconds == "" {
		return time.Time{}, nil
	}
	rangeSeconds, err := strconv.ParseUint(rawRangeSeconds, 10, 64)
	if rangeSeconds <= 0 {
		return time.Time{}, nil
	}
	if err != nil {
		return time.Time{}, err
	}
	fromTime := time.Now().Add(-time.Second * time.Duration(rangeSeconds)).UTC()
	return fromTime, nil
}

func (s *Server) GetLeaderboardStats(w http.ResponseWriter, r *http.Request) {
	rawRangeSeconds := r.URL.Query().Get("range")
	fromTime, err := getRangeStartTime(rawRangeSeconds)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	stats, err := s.db.GetLeaderboardStats(fromTime)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res := &LeaderboardStatsResponse{Stats: stats}
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResponse); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
