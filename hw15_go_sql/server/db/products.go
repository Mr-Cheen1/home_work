package db

// Функции для работы с товарами.
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
		if err = rows.Scan(&id, &name, &price); err != nil {
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
