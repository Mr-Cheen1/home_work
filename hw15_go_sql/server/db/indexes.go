package db

// Создание индексов.
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
