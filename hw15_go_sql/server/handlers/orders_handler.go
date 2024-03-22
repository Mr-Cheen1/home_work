package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/db"
)

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		handleGetOrders(w, r)
	case http.MethodPost:
		handleCreateOrder(w, r)
	case http.MethodDelete:
		handleDeleteOrder(w, r)
	}
}

func handleGetOrders(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	orders, err := db.GetOrdersByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"path":   r.URL.Path,
		"params": r.URL.Query(),
		"data":   orders,
		"status": http.StatusOK,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var order map[string]interface{}
	json.NewDecoder(r.Body).Decode(&order)

	userID, _ := strconv.Atoi(order["user_id"].(string))
	totalAmount, _ := order["total_amount"].(float64)

	orderID, err := db.InsertOrder(userID, totalAmount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"path":   r.URL.Path,
		"params": r.URL.Query(),
		"data":   map[string]int{"order_id": orderID},
		"status": http.StatusCreated,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func handleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	err := db.DeleteOrder(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"path":   r.URL.Path,
		"params": r.URL.Query(),
		"status": http.StatusNoContent,
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(response)
}
