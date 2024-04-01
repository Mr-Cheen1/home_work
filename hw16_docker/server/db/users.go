package db

import "database/sql"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func InsertUser(user User) (int, error) {
	var userID int
	err := DB.QueryRow(
		"INSERT INTO Users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Name, user.Email, user.Password,
	).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func UpdateUser(user User) error {
	_, err := DB.Exec(
		"UPDATE Users SET name = $1, email = $2, password = $3 WHERE id = $4",
		user.Name, user.Email, user.Password, user.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	_, err := DB.Exec("DELETE FROM Users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(userID ...int) ([]User, error) {
	query := "SELECT * FROM Users"
	if len(userID) > 0 {
		query += " WHERE id = $1"
		return queryRows(query, func(rows *sql.Rows) (User, error) {
			var user User
			err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
			return user, err
		}, userID[0])
	}
	return queryRows(query, func(rows *sql.Rows) (User, error) {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		return user, err
	})
}
