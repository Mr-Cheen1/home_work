package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/db"
)

const ErrMethodNotAllowed = "Method not allowed"

func AddProductToOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	response := map[string]interface{}{
		"path":   r.URL.Path,
		"params": r.URL.Query(),
	}

	if r.Method != http.MethodPost {
		response["status"] = http.StatusMethodNotAllowed
		response["error"] = ErrMethodNotAllowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	var data map[string]int
	json.NewDecoder(r.Body).Decode(&data)
	orderID := data["order_id"]
	productID := data["product_id"]

	err := db.AddProductToOrder(orderID, productID)
	if err != nil {
		response["status"] = http.StatusInternalServerError
		response["error"] = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	response["status"] = http.StatusCreated
	response["data"] = map[string]interface{}{
		"order_id":   orderID,
		"product_id": productID,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
