package db

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func InsertProduct(product *Product) error {
	var lastInsertID int
	err := DB.QueryRow(
		"INSERT INTO Products (name, price) VALUES ($1, $2) RETURNING id",
		product.Name, product.Price,
	).Scan(&lastInsertID)
	if err != nil {
		return err
	}
	product.ID = lastInsertID
	return nil
}

func UpdateProduct(product Product) error {
	_, err := DB.Exec("UPDATE Products SET name = $1, price = $2 WHERE id = $3", product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id int) error {
	_, err := DB.Exec("DELETE FROM Products WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func GetProducts() ([]Product, error) {
	rows, err := DB.Query("SELECT * FROM Products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
