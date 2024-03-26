package db

import "database/sql"

// Функции для работы с отношением заказов и товаров.
func AddProductToOrder(db *sql.DB, orderID, productID, quantity int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(
		"INSERT INTO OrderProducts (order_id, product_id, quantity) VALUES ($1, $2, $3)",
		orderID, productID, quantity,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		"UPDATE Orders SET total_amount = total_amount + (SELECT price FROM Products WHERE id = $1) * $2 WHERE id = $3",
		productID, quantity, orderID,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func RemoveProductFromOrder(db *sql.DB, orderID, productID, quantity int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(
		"UPDATE Orders SET total_amount = total_amount - (SELECT price FROM Products WHERE id = $1) * $2 WHERE id = $3",
		productID, quantity, orderID,
	)
	if err != nil {
		return err
	}

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
