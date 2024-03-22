package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/db"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	response := map[string]interface{}{
		"path":   r.URL.Path,
		"params": r.URL.Query(),
	}

	switch r.Method {
	case http.MethodGet:
		getProducts(w, response)
	case http.MethodPost:
		createProduct(w, r, response)
	case http.MethodPut:
		updateProduct(w, r, response)
	case http.MethodDelete:
		deleteProduct(w, r, response)
	}
}

func getProducts(w http.ResponseWriter, response map[string]interface{}) {
	products, err := db.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response["data"] = products
	response["status"] = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func createProduct(w http.ResponseWriter, r *http.Request, response map[string]interface{}) {
	var product map[string]interface{}
	json.NewDecoder(r.Body).Decode(&product)
	err := db.InsertProduct(product["name"].(string), product["price"].(float64))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response["data"] = product
	response["status"] = http.StatusCreated
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func updateProduct(w http.ResponseWriter, r *http.Request, response map[string]interface{}) {
	var product map[string]interface{}
	json.NewDecoder(r.Body).Decode(&product)
	id, _ := strconv.Atoi(product["id"].(string))
	err := db.UpdateProduct(id, product["name"].(string), product["price"].(float64))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response["data"] = product
	response["status"] = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func deleteProduct(w http.ResponseWriter, r *http.Request, response map[string]interface{}) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	err := db.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response["status"] = http.StatusNoContent
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(response)
}
