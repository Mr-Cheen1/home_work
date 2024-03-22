package db

import (
	"database/sql"

	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL для использования с database/sql.
)

var db *sql.DB

func InitDB() {
	// Инициализация подключения к БД.
	var err error
	db, err = sql.Open("postgres", "postgres://store_user:123456@localhost:8080/online_store?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	db.Close()
}
