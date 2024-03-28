package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsertOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	DB = db

	order := Order{
		UserID:      1,
		OrderDate:   "2023-06-08",
		TotalAmount: 100.0,
	}

	mock.ExpectQuery("INSERT INTO Orders").
		WithArgs(order.UserID, order.OrderDate, order.TotalAmount).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	orderID, err := InsertOrder(order)
	assert.NoError(t, err)
	assert.Equal(t, 1, orderID)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	DB = db

	orderID := 1

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM OrderProducts").
		WithArgs(orderID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("DELETE FROM Orders").
		WithArgs(orderID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = DeleteOrder(orderID)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetOrders(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	DB = db

	userID := 1

	rows := sqlmock.NewRows([]string{"id", "user_id", "order_date", "total_amount"}).
		AddRow(1, userID, "2023-06-08", 100.0).
		AddRow(2, userID, "2023-06-09", 200.0)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT \\* FROM Orders WHERE user_id = \\$1").
		WithArgs(userID).
		WillReturnRows(rows)
	mock.ExpectCommit()

	orders, err := GetOrders(userID)
	assert.NoError(t, err)
	assert.Len(t, orders, 2)

	assert.NoError(t, mock.ExpectationsWereMet())
}
