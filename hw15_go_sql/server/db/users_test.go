package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	DB = db

	user := User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password",
	}

	mock.ExpectQuery("INSERT INTO Users").
		WithArgs(user.Name, user.Email, user.Password).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	userID, err := InsertUser(user)
	assert.NoError(t, err)
	assert.Equal(t, 1, userID)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	DB = db

	user := User{
		ID:       1,
		Name:     "John Updated",
		Email:    "john.updated@example.com",
		Password: "newpassword",
	}

	mock.ExpectExec("UPDATE Users SET").
		WithArgs(user.Name, user.Email, user.Password, user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = UpdateUser(user)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	DB = db

	userID := 1

	mock.ExpectExec("DELETE FROM Users WHERE").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = DeleteUser(userID)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	DB = db

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(1, "John Doe", "john@example.com", "password").
		AddRow(2, "Jane Smith", "jane@example.com", "password")

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT \\* FROM Users").WillReturnRows(rows)
	mock.ExpectCommit()

	users, err := GetUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 2)

	assert.NoError(t, mock.ExpectationsWereMet())
}
