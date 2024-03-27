package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/db"
	"github.com/Mr-Cheen1/home_work/hw15_go_sql/server/handlers"
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

	// Создание экземпляра сервера.
	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%s", address, port),
		Handler:           nil,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Регистрация обработчиков.
	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/products", handlers.ProductsHandler)
	http.HandleFunc("/orders", handlers.OrdersHandler)
	http.HandleFunc("/stats", handlers.UserStatsHandler)
	http.HandleFunc("/order-products/add", handlers.AddProductToOrderHandler)
	http.HandleFunc("/order-products/remove", handlers.RemoveProductFromOrderHandler)

	// Запуск сервера в отдельной горутине.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	log.Printf("Server listening on %s:%s", address, port)

	// Ожидание сигнала для graceful shutdown.
	if err := gracefulShutdown(srv); err != nil {
		log.Println("Failed to gracefully shutdown:", err)
		return
	}
}

func gracefulShutdown(srv *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown:", err)
		return err
	}
	log.Println("Server exiting")
	return nil
}
