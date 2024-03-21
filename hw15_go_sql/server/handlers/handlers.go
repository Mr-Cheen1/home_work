package handlers

import (
	"encoding/json"
	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/db"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	response := map[string]interface{}{
		"path":   r.URL.Path,
		"params": r.URL.Query(),
	}

	switch r.Method {
	case http.MethodGet:
		users, err := db.GetUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response["data"] = users
		response["status"] = http.StatusOK
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	case http.MethodPost:
		var user map[string]string
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			response["status"] = http.StatusBadRequest
			response["error"] = "Invalid request payload"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Удаление поля "id" из данных JSON, если оно присутствует
		delete(user, "id")

		err = db.InsertUser(user["name"], user["email"], user["password"])
		if err != nil {
			if strings.Contains(err.Error(), "повторяющееся значение ключа нарушает ограничение уникальности") {
				response["status"] = http.StatusConflict
				response["error"] = "Email уже существует"
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
	case http.MethodPut:
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
	case http.MethodDelete:
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
		response["status"] = http.StatusNoContent
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(response)
	default:
		response["status"] = http.StatusMethodNotAllowed
		response["error"] = "Method not allowed"
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
	}
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	response := map[string]interface{}{
		"path":   r.URL.Path,
		"params": r.URL.Query(),
	}

	switch r.Method {
	case http.MethodGet:
		products, err := db.GetProducts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response["data"] = products
		response["status"] = http.StatusOK
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	case http.MethodPost:
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
	case http.MethodPut:
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
	case http.MethodDelete:
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
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	response := map[string]interface{}{
		"path":   r.URL.Path,
		"params": r.URL.Query(),
	}

	switch r.Method {
	case http.MethodGet:
		userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
		orders, err := db.GetOrdersByUser(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response["data"] = orders
		response["status"] = http.StatusOK
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	case http.MethodPost:
		var order map[string]interface{}
		json.NewDecoder(r.Body).Decode(&order)
		userID, _ := strconv.Atoi(order["user_id"].(string))
		totalAmount, _ := order["total_amount"].(float64)
		orderID, err := db.InsertOrder(userID, totalAmount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response["data"] = map[string]int{"order_id": orderID}
		response["status"] = http.StatusCreated
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	case http.MethodDelete:
		orderID, _ := strconv.Atoi(r.URL.Query().Get("id"))
		err := db.DeleteOrder(orderID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response["status"] = http.StatusNoContent
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(response)
	}
}

func UserStatsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	response := map[string]interface{}{
		"path":   r.URL.Path,
		"params": r.URL.Query(),
	}

	if r.Method != http.MethodGet {
		response["status"] = http.StatusMethodNotAllowed
		response["error"] = "Method not allowed"
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	stats, err := db.GetUserStats(userID)
	if err != nil {
		response["status"] = http.StatusInternalServerError
		response["error"] = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	response["data"] = stats
	response["status"] = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func AddProductToOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	response := map[string]interface{}{
		"path":   r.URL.Path,
		"params": r.URL.Query(),
	}

	if r.Method != http.MethodPost {
		response["status"] = http.StatusMethodNotAllowed
		response["error"] = "Method not allowed"
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
