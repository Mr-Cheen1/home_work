package db

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

const (
	avgPriceQuery         = "COALESCE\\(AVG\\(p.price\\), 0\\) AS avg_product_price "
	fromQuery             = "FROM Orders o "
	leftJoinOrderProducts = "LEFT JOIN OrderProducts op ON o.id = op.order_id "
	leftJoinProducts      = "LEFT JOIN Products p ON op.product_id = p.id "
	selectQuery           = "SELECT COUNT\\(\\*\\) AS total_orders, "
	totalAmountQuery      = "COALESCE\\(SUM\\(o.total_amount\\), 0\\) AS total_amount, "
	whereUserID           = "WHERE o.user_id = \\$1"
)

func TestGetUserStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	DB = db

	userID := 1

	// Тест для случая, когда userID передан.
	rows := sqlmock.NewRows([]string{"total_orders", "total_amount", "avg_product_price"}).
		AddRow(5, 1000.0, 50.0)

	mock.ExpectBegin()
	mock.ExpectQuery(selectQuery + totalAmountQuery + avgPriceQuery +
		fromQuery + leftJoinOrderProducts + leftJoinProducts + whereUserID).
		WithArgs(userID).WillReturnRows(rows)
	mock.ExpectCommit()

	stats, err := GetUserStats(&userID)
	assert.NoError(t, err)
	assert.Equal(t, &UserStats{TotalOrders: 5, TotalAmount: 1000.0, AvgProductPrice: 50.0}, stats)

	// Тест для случая, когда userID не передан.
	rows = sqlmock.NewRows([]string{"total_orders", "total_amount", "avg_product_price"}).
		AddRow(10, 2000.0, 75.0)

	mock.ExpectBegin()
	mock.ExpectQuery(selectQuery + totalAmountQuery + avgPriceQuery +
		fromQuery + leftJoinOrderProducts + leftJoinProducts).
		WillReturnRows(rows)
	mock.ExpectCommit()

	stats, err = GetUserStats(nil)
	assert.NoError(t, err)
	assert.Equal(t, &UserStats{TotalOrders: 10, TotalAmount: 2000.0, AvgProductPrice: 75.0}, stats)

	// Тест для случая, когда нет данных.
	mock.ExpectBegin()
	mock.ExpectQuery(selectQuery + totalAmountQuery + avgPriceQuery +
		fromQuery + leftJoinOrderProducts + leftJoinProducts + whereUserID).
		WithArgs(userID).WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()

	stats, err = GetUserStats(&userID)
	assert.NoError(t, err)
	assert.Equal(t, &UserStats{}, stats)

	// Тест для случая, когда возникает ошибка при выполнении запроса.
	mock.ExpectBegin()
	mock.ExpectQuery(selectQuery + totalAmountQuery + avgPriceQuery +
		fromQuery + leftJoinOrderProducts + leftJoinProducts + whereUserID).
		WithArgs(userID).WillReturnError(errors.New("some error"))
	mock.ExpectRollback()

	stats, err = GetUserStats(&userID)
	assert.Error(t, err)
	assert.Nil(t, stats)

	// Тест для случая, когда возникает ошибка при коммите транзакции.
	rows = sqlmock.NewRows([]string{"total_orders", "total_amount", "avg_product_price"}).
		AddRow(5, 1000.0, 50.0)

	mock.ExpectBegin()
	mock.ExpectQuery(selectQuery + totalAmountQuery + avgPriceQuery +
		fromQuery + leftJoinOrderProducts + leftJoinProducts + whereUserID).
		WithArgs(userID).WillReturnRows(rows)
	mock.ExpectCommit().WillReturnError(errors.New("commit error"))

	stats, err = GetUserStats(&userID)
	assert.Error(t, err)
	assert.Nil(t, stats)

	// Проверяем, что все ожидаемые запросы были выполнены.
	assert.NoError(t, mock.ExpectationsWereMet())
}
