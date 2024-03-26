package db

import "database/sql"

type Order struct {
	ID          int     `json:"id"`
	UserID      int     `json:"userId"`
	OrderDate   string  `json:"orderDate"`
	TotalAmount float64 `json:"totalAmount"`
}

func InsertOrder(order Order) (int, error) {
	var orderID int
	err := DB.QueryRow(
		"INSERT INTO Orders (user_id, order_date, total_amount) VALUES ($1, $2, $3) RETURNING id",
		order.UserID, order.OrderDate, order.TotalAmount,
	).Scan(&orderID)
	if err != nil {
		return 0, err
	}
	return orderID, nil
}

func DeleteOrder(orderID int) error {
	tx, err := DB.Begin()
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

func GetOrders(userID ...int) ([]Order, error) {
	query := "SELECT * FROM Orders"
	if len(userID) > 0 {
		query += " WHERE user_id = $1"
		return queryRows(query, func(rows *sql.Rows) (Order, error) {
			var order Order
			err := rows.Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount)
			return order, err
		}, userID[0])
	}
	return queryRows(query, func(rows *sql.Rows) (Order, error) {
		var order Order
		err := rows.Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount)
		return order, err
	})
}
