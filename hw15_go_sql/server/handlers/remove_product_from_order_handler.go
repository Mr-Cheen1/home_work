package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/db"
)

type RemoveProductFromOrderData struct {
	Path      string     `json:"path"`
	Params    url.Values `json:"params"`
	Status    int        `json:"status,omitempty"`
	Error     string     `json:"error,omitempty"`
	OrderID   int        `json:"orderId"`
	ProductID int        `json:"productId"`
	Quantity  int        `json:"quantity"`
}

func RemoveProductFromOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	data := RemoveProductFromOrderData{
		Path:   r.URL.Path,
		Params: r.URL.Query(),
	}
	if r.Method != http.MethodDelete {
		data.Status = http.StatusMethodNotAllowed
		data.Error = "Method not allowed"
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(data)
		return
	}

	orderID, _ := strconv.Atoi(r.URL.Query().Get("order_id"))
	productID, _ := strconv.Atoi(r.URL.Query().Get("product_id"))
	quantity, _ := strconv.Atoi(r.URL.Query().Get("quantity"))
	data.OrderID = orderID
	data.ProductID = productID
	data.Quantity = quantity

	err := db.RemoveProductFromOrder(db.DB, data.OrderID, data.ProductID, data.Quantity)
	if err != nil {
		data.Status = http.StatusInternalServerError
		data.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(data)
		return
	}
	data.Status = http.StatusNoContent
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(data)
}
