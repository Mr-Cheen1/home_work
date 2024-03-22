package db

// Функции для работы с заказами.
func InsertOrder(userID int, totalAmount float64) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var orderID int
	err = tx.QueryRow(
		"INSERT INTO Orders (user_id, order_date, total_amount) VALUES ($1, NOW(), $2) RETURNING id",
		userID, totalAmount,
	).Scan(&orderID)
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
		var id int
		var orderDate string
		var totalAmount float64
		err = rows.Scan(&id, &userID, &orderDate, &totalAmount)
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
