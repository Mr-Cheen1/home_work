package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/db"
)

type OrderResponse struct {
	Path    string      `json:"path"`
	Params  string      `json:"params"`
	Data    interface{} `json:"data,omitempty"`
	Status  int         `json:"status"`
	OrderID int         `json:"orderId,omitempty"`
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	response := OrderResponse{
		Path:   r.URL.Path,
		Params: r.URL.Query().Encode(),
	}

	switch r.Method {
	case http.MethodGet:
		getOrders(w, r, &response)
	case http.MethodPut:
		createOrder(w, r, &response)
	case http.MethodDelete:
		deleteOrder(w, r, &response)
	}
}

func getOrders(w http.ResponseWriter, r *http.Request, response *OrderResponse) {
	userIDStr := r.URL.Query().Get("user_id")

	var orders []db.Order
	var err error

	if userIDStr != "" {
		userID, _ := strconv.Atoi(userIDStr)
		orders, err = db.GetOrders(userID)
	} else {
		orders, err = db.GetOrders()
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Data = orders
	response.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func createOrder(w http.ResponseWriter, r *http.Request, response *OrderResponse) {
	var order db.Order
	json.NewDecoder(r.Body).Decode(&order)
	orderID, err := db.InsertOrder(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.OrderID = orderID
	response.Status = http.StatusCreated
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func deleteOrder(w http.ResponseWriter, r *http.Request, response *OrderResponse) {
	orderID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	err := db.DeleteOrder(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Status = http.StatusNoContent
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(response)
}
