package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddProductToOrder(t *testing.T) {
	// Создаем mock-соединение с базой данных.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Настраиваем ожидаемые запросы к базе данных.
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO OrderProducts").WithArgs(1, 1, 2).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE Orders").WithArgs(1, 2, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Вызываем тестируемую функцию с mock-соединением.
	err = AddProductToOrder(db, 1, 1, 2)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// Проверяем, что все ожидаемые запросы были выполнены.
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRemoveProductFromOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE Orders").WithArgs(1, 2, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("DELETE FROM OrderProducts").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = RemoveProductFromOrder(db, 1, 1, 2)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
