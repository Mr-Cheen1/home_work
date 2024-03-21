package main

import (
	"fmt"
	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/db"
	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run server/main.go <address> <port>")
		os.Exit(1)
	}

	address := os.Args[1]
	port := os.Args[2]

	db.InitDB()
	defer db.CloseDB()

	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/products", handlers.ProductsHandler)
	http.HandleFunc("/orders", handlers.OrdersHandler)
	http.HandleFunc("/stats", handlers.UserStatsHandler)
	http.HandleFunc("/order-products/add", handlers.AddProductToOrderHandler)
	http.HandleFunc("/order-products/remove", handlers.RemoveProductFromOrderHandler)

	log.Printf("Server listening on %s:%s", address, port)
	log.Fatal(http.ListenAndServe(address+":"+port, nil))
}
