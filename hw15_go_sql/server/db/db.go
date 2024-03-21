package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	// Инициализация подключения к БД
	var err error
	db, err = sql.Open("postgres", "postgres://store_user:123456@localhost:8080/online_store?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	db.Close()
}

// Функции для работы с пользователями
func InsertUser(name, email, password string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO Users (name, email, password) VALUES ($1, $2, $3)", name, email, password)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(id int, name, email, password string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE Users SET name = $1, email = $2, password = $3 WHERE id = $4", name, email, password, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM Users WHERE id = $1", id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GetUsers() ([]map[string]interface{}, error) {
	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var name, email, password string
		err := rows.Scan(&id, &name, &email, &password)
		if err != nil {
			return nil, err
		}
		user := map[string]interface{}{
			"id":       id,
			"name":     name,
			"email":    email,
			"password": password,
		}
		users = append(users, user)
	}
	return users, nil
}

// Функции для работы с товарами
func InsertProduct(name string, price float64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO Products (name, price) VALUES ($1, $2)", name, price)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(id int, name string, price float64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE Products SET name = $1, price = $2 WHERE id = $3", name, price, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(id int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM Products WHERE id = $1", id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GetProducts() ([]map[string]interface{}, error) {
	rows, err := db.Query("SELECT * FROM Products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var name string
		var price float64
		err := rows.Scan(&id, &name, &price)
		if err != nil {
			return nil, err
		}
		product := map[string]interface{}{
			"id":    id,
			"name":  name,
			"price": price,
		}
		products = append(products, product)
	}
	return products, nil
}

// Функции для работы с заказами
func InsertOrder(userID int, totalAmount float64) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var orderID int
	err = tx.QueryRow("INSERT INTO Orders (user_id, order_date, total_amount) VALUES ($1, NOW(), $2) RETURNING id", userID, totalAmount).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func DeleteOrder(orderID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM OrderProducts WHERE order_id = $1", orderID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM Orders WHERE id = $1", orderID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GetOrdersByUser(userID int) ([]map[string]interface{}, error) {
	rows, err := db.Query("SELECT * FROM Orders WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []map[string]interface{}{}
	for rows.Next() {
		var id, userID int
		var orderDate string
		var totalAmount float64
		err := rows.Scan(&id, &userID, &orderDate, &totalAmount)
		if err != nil {
			return nil, err
		}
		order := map[string]interface{}{
			"id":           id,
			"user_id":      userID,
			"order_date":   orderDate,
			"total_amount": totalAmount,
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func GetUserStats(userID int) (map[string]interface{}, error) {
	var totalOrders int
	var totalAmount, avgProductPrice float64

	err := db.QueryRow("SELECT COUNT(*) FROM Orders WHERE user_id = $1", userID).Scan(&totalOrders)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow("SELECT SUM(total_amount) FROM Orders WHERE user_id = $1", userID).Scan(&totalAmount)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow(`
		 SELECT AVG(p.price) 
		 FROM Products p
		 JOIN OrderProducts op ON p.id = op.product_id
		 JOIN Orders o ON op.order_id = o.id
		 WHERE o.user_id = $1
	`, userID).Scan(&avgProductPrice)
	if err != nil {
		return nil, err
	}

	stats := map[string]interface{}{
		"total_orders":      totalOrders,
		"total_amount":      totalAmount,
		"avg_product_price": avgProductPrice,
	}
	return stats, nil
}

// Функции для работы с взимодействиями заказов и товаров
func AddProductToOrder(orderID, productID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO OrderProducts (order_id, product_id) VALUES ($1, $2)", orderID, productID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func RemoveProductFromOrder(orderID, productID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM OrderProducts WHERE order_id = $1 AND product_id = $2", orderID, productID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Создание индексов
func CreateIndexes() error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("CREATE INDEX idx_orders_user_id ON Orders (user_id)")
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE INDEX idx_orderproducts_order_id ON OrderProducts (order_id)")
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE INDEX idx_orderproducts_product_id ON OrderProducts (product_id)")
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
