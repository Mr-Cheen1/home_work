package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/db"
)

type UserStatsData struct {
	Path   string `json:"path"`
	Params struct {
		UserID int `json:"userId"`
	} `json:"params"`
	Status int           `json:"status"`
	Error  string        `json:"error,omitempty"`
	Stats  *db.UserStats `json:"stats,omitempty"`
}

func UserStatsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	data := UserStatsData{
		Path: r.URL.Path,
	}

	if r.Method != http.MethodGet {
		data.Status = http.StatusMethodNotAllowed
		data.Error = "Method not allowed"
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(data)
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	var userID *int

	if userIDStr != "" {
		parsedID, err := strconv.Atoi(userIDStr)
		if err != nil {
			data.Status = http.StatusBadRequest
			data.Error = "user_id must be an integer"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(data)
			return
		}
		userID = &parsedID
		data.Params.UserID = parsedID
	}

	stats, err := db.GetUserStats(userID)
	if err != nil {
		data.Status = http.StatusInternalServerError
		data.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(data)
		return
	}

	data.Stats = stats
	data.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
