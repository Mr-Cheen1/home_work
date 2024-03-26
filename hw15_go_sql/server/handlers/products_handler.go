package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/db"
)

type ProductResponse struct {
	Path   string      `json:"path"`
	Params string      `json:"params"`
	Data   interface{} `json:"data,omitempty"`
	Status int         `json:"status"`
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	response := ProductResponse{
		Path:   r.URL.Path,
		Params: r.URL.Query().Encode(),
	}

	switch r.Method {
	case http.MethodGet:
		getProducts(w, r, &response)
	case http.MethodPut:
		createProduct(w, r, &response)
	case http.MethodPost:
		updateProduct(w, r, &response)
	case http.MethodDelete:
		deleteProduct(w, r, &response)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request, response *ProductResponse) {
	idStr := r.URL.Query().Get("id")
	products, err := db.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if idStr != "" {
		id, parseErr := strconv.Atoi(idStr)
		if parseErr != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}
		for _, product := range products {
			if product.ID == id {
				response.Data = []db.Product{product}
				break
			}
		}
	} else {
		response.Data = products
	}

	response.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func createProduct(w http.ResponseWriter, r *http.Request, response *ProductResponse) {
	var product db.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := db.InsertProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Data = product
	response.Status = http.StatusCreated
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func updateProduct(w http.ResponseWriter, r *http.Request, response *ProductResponse) {
	var product db.Product
	json.NewDecoder(r.Body).Decode(&product)
	err := db.UpdateProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Data = product
	response.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func deleteProduct(w http.ResponseWriter, r *http.Request, response *ProductResponse) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	err := db.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Status = http.StatusNoContent
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(response)
}
