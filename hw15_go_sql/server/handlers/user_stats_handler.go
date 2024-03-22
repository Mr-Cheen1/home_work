package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/db"
)

func UserStatsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	response := map[string]interface{}{
		"path":   r.URL.Path,
		"params": r.URL.Query(),
	}

	if r.Method != http.MethodGet {
		response["status"] = http.StatusMethodNotAllowed
		response["error"] = ErrMethodNotAllowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	stats, err := db.GetUserStats(userID)
	if err != nil {
		response["status"] = http.StatusInternalServerError
		response["error"] = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	response["data"] = stats
	response["status"] = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
