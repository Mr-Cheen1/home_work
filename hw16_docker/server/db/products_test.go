package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsertProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	DB = db

	product := &Product{
		Name:  "Test Product",
		Price: 9.99,
	}

	mock.ExpectQuery("INSERT INTO Products").
		WithArgs(product.Name, product.Price).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	err = InsertProduct(product)
	assert.NoError(t, err)
	assert.Equal(t, 1, product.ID)
}

func TestUpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	DB = db

	product := Product{
		ID:    1,
		Name:  "Updated Product",
		Price: 19.99,
	}

	mock.ExpectExec("UPDATE Products").
		WithArgs(product.Name, product.Price, product.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = UpdateProduct(product)
	assert.NoError(t, err)
}

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	DB = db

	productID := 1

	mock.ExpectExec("DELETE FROM Products").
		WithArgs(productID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = DeleteProduct(productID)
	assert.NoError(t, err)
}

func TestGetProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	DB = db

	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(1, "Product 1", 9.99).
		AddRow(2, "Product 2", 19.99)

	mock.ExpectQuery("SELECT \\* FROM Products").
		WillReturnRows(rows)

	products, err := GetProducts()
	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, 19.99, products[1].Price)
}
