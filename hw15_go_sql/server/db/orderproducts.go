package db

// Функции для работы с отношением заказов и товаров.
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
