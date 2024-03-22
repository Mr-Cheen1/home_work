package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/db"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	response := map[string]interface{}{
		"path":   r.URL.Path,
		"params": r.URL.Query(),
	}

	switch r.Method {
	case http.MethodGet:
		handleGetUsers(w, response)
	case http.MethodPost:
		handleCreateUser(w, r, response)
	case http.MethodPut:
		handleUpdateUser(w, r, response)
	case http.MethodDelete:
		handleDeleteUser(w, r, response)
	default:
		handleMethodNotAllowed(w, response)
	}
}

func handleGetUsers(w http.ResponseWriter, response map[string]interface{}) {
	users, err := db.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response["data"] = users
	response["status"] = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request, response map[string]interface{}) {
	var user map[string]string
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response["status"] = http.StatusBadRequest
		response["error"] = "Invalid request payload"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = db.InsertUser(user["name"], user["email"], user["password"])
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			response["status"] = http.StatusConflict
			response["error"] = "Email already exists"
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(response)
			return
		}
		response["status"] = http.StatusInternalServerError
		response["error"] = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	response["data"] = user
	response["status"] = http.StatusCreated
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request, response map[string]interface{}) {
	var user map[string]string
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response["status"] = http.StatusBadRequest
		response["error"] = "Invalid request payload"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	id, err := strconv.Atoi(user["id"])
	if err != nil {
		response["status"] = http.StatusBadRequest
		response["error"] = "Invalid user ID"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = db.UpdateUser(id, user["name"], user["email"], user["password"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response["data"] = user
	response["status"] = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request, response map[string]interface{}) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response["status"] = http.StatusBadRequest
		response["error"] = "Invalid user ID"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = db.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response["status"] = http.StatusOK
	response["message"] = "User successfully deleted"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handleMethodNotAllowed(w http.ResponseWriter, response map[string]interface{}) {
	response["status"] = http.StatusMethodNotAllowed
	response["error"] = ErrMethodNotAllowed
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(response)
}
