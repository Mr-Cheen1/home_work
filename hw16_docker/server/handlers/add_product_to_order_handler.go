package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/Mr-Cheen1/home_work/hw16_docker/server/db"
)

const ErrMethodNotAllowed = "Method not allowed"

type AddProductToOrderData struct {
	Path      string     `json:"path"`
	Params    url.Values `json:"params"`
	Status    int        `json:"status,omitempty"`
	Error     string     `json:"error,omitempty"`
	OrderID   int        `json:"orderId"`
	ProductID int        `json:"productId"`
	Quantity  int        `json:"quantity"`
}

func AddProductToOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	data := AddProductToOrderData{
		Path:   r.URL.Path,
		Params: r.URL.Query(),
	}
	if r.Method != http.MethodPut {
		data.Status = http.StatusMethodNotAllowed
		data.Error = ErrMethodNotAllowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(data)
		return
	}

	json.NewDecoder(r.Body).Decode(&data)
	data.Path = r.URL.Path // Устанавливаем значение поля Path

	err := db.AddProductToOrder(db.DB, data.OrderID, data.ProductID, data.Quantity)
	if err != nil {
		data.Status = http.StatusInternalServerError
		data.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(data)
		return
	}
	data.Status = http.StatusCreated
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
