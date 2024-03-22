package db

// Функции для работы с пользователями.
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
		err = rows.Scan(&id, &name, &email, &password)
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
