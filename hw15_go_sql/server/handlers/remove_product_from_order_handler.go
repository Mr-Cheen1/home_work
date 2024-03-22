package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/db"
)

func RemoveProductFromOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	response := map[string]interface{}{
		"path":   r.URL.Path,
		"params": r.URL.Query(),
	}

	if r.Method != http.MethodDelete {
		response["status"] = http.StatusMethodNotAllowed
		response["error"] = "Method not allowed"
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	orderID, _ := strconv.Atoi(r.URL.Query().Get("order_id"))
	productID, _ := strconv.Atoi(r.URL.Query().Get("product_id"))

	err := db.RemoveProductFromOrder(orderID, productID)
	if err != nil {
		response["status"] = http.StatusInternalServerError
		response["error"] = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	response["status"] = http.StatusNoContent
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(response)
}
