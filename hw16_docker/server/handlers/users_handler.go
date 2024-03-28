package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Mr-Cheen1/home_work/hw16_docker/server/db"
)

type UserResponse struct {
	Path   string      `json:"path"`
	Params string      `json:"params"`
	Data   interface{} `json:"data,omitempty"`
	Status int         `json:"status"`
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	response := UserResponse{
		Path:   r.URL.Path,
		Params: r.URL.Query().Encode(),
	}

	switch r.Method {
	case http.MethodGet:
		getUsers(w, r, &response)
	case http.MethodPut:
		createUser(w, r, &response)
	case http.MethodPost:
		updateUser(w, r, &response)
	case http.MethodDelete:
		deleteUser(w, r, &response)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request, response *UserResponse) {
	userIDStr := r.URL.Query().Get("user_id")

	var users []db.User
	var err error

	if userIDStr != "" {
		userID, _ := strconv.Atoi(userIDStr)
		users, err = db.GetUsers(userID)
	} else {
		users, err = db.GetUsers()
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Data = users
	response.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func createUser(w http.ResponseWriter, r *http.Request, response *UserResponse) {
	var user db.User
	json.NewDecoder(r.Body).Decode(&user)
	insertedUserID, err := db.InsertUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.ID = insertedUserID
	response.Data = user
	response.Status = http.StatusCreated
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func updateUser(w http.ResponseWriter, r *http.Request, response *UserResponse) {
	var user db.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error decoding user data: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.ID == 0 {
		log.Printf("User ID is required")
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	err := db.UpdateUser(user)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Data = user
	response.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func deleteUser(w http.ResponseWriter, r *http.Request, response *UserResponse) {
	id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	err = db.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Status = http.StatusNoContent
	w.WriteHeader(http.StatusNoContent)
}
