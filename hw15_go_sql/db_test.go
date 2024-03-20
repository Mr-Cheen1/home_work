package main

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// Вспомогательная функция для создания мока базы данных и настройки ожидаемых действий.
func setupMock() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, fmt.Errorf("error when creating a database moc: %w", err)
	}
	return db, mock, nil
}

// Вспомогательная функция для создания временного файла с SQL запросом.
func createTempSQLFile(query string) (string, error) {
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		return "", fmt.Errorf("error when creating a temporary file: %w", err)
	}
	defer tmpfile.Close()

	defer tmpfile.Close()

	_, err = tmpfile.Write([]byte(query))
	if err != nil {
		return "", fmt.Errorf("error when writing to a temporary file: %w", err)
	}

	return tmpfile.Name(), nil
}

func TestExecuteSQLFromFile_SelectQuery_MockNotImplemented(t *testing.T) {
	db, mock, err := setupMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	query := "SELECT id, name FROM users"
	filePath, err := createTempSQLFile(query)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filePath)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "John Doe").
		AddRow(2, "Jane Doe")

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectCommit()

	database := &Database{conn: db}

	result, err := database.ExecuteSQLFromFile(filePath)
	if err != nil {
		t.Errorf("Error was not expected while executing SQL from file: %s", err)
	}

	expected := []map[string]interface{}{
		{"id": int64(1), "name": "John Doe"},
		{"id": int64(2), "name": "Jane Doe"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result %v, got %v", expected, result)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Not all expected actions were completed: %s", err)
	}
}

func TestCleanAndDetectType(t *testing.T) {
	tests := []struct {
		name      string
		query     string
		wantClean string
		wantType  string
	}{
		{
			name:      "SELECT query",
			query:     "SELECT * FROM users -- Get all users",
			wantClean: "SELECT * FROM users ",
			wantType:  "SELECT",
		},
		{
			name:      "INSERT query",
			query:     "INSERT INTO users (name) VALUES ('John Doe')",
			wantClean: "INSERT INTO users (name) VALUES ('John Doe') ",
			wantType:  "INSERT",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClean, gotType := cleanAndDetectType(tt.query)
			if gotClean != tt.wantClean || gotType != tt.wantType {
				t.Errorf("cleanAndDetectType() gotClean = %v, want %v, "+
					"gotType = %v, want %v", gotClean, tt.wantClean, gotType, tt.wantType)
			}
		})
	}
}

func TestNewDatabase(t *testing.T) {
	connString := "postgres://user:password@localhost/dbname?sslmode=disable"
	db, err := NewDatabase(connString)
	if err != nil {
		t.Fatalf("Failed to create database: %v", err)
	}
	if db.conn == nil {
		t.Error("Expected database connection to be non-nil")
	}
}

func TestExecuteSQLFromFile_InsertQuery(t *testing.T) {
	db, mock, err := setupMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	query := "INSERT INTO users (name) VALUES ('John Doe')"
	filePath, err := createTempSQLFile(query)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filePath)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	database := &Database{conn: db}

	_, err = database.ExecuteSQLFromFile(filePath)
	if err != nil {
		t.Errorf("Error was not expected while executing SQL from file: %s", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Not all expected actions were completed: %s", err)
	}
}

func TestReadAndCleanQueryFromFile(t *testing.T) {
	query := "SELECT * FROM users -- This is a comment"
	filePath, err := createTempSQLFile(query)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filePath)

	db := &Database{}
	cleanedQuery, queryType, err := db.readAndCleanQueryFromFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read and clean query from file: %v", err)
	}

	expectedCleanedQuery := "SELECT * FROM users "
	if cleanedQuery != expectedCleanedQuery || queryType != "SELECT" {
		t.Errorf(
			"Expected cleaned query to be '%s' and query type to be 'SELECT', got '%s' and '%s'",
			expectedCleanedQuery, cleanedQuery, queryType,
		)
	}
}

func TestExecuteSQLFromFile_CreateQuery(t *testing.T) {
	db, mock, err := setupMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	query := "CREATE TABLE test_table (id SERIAL PRIMARY KEY, name VARCHAR(100))"
	filePath, err := createTempSQLFile(query)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filePath)

	mock.ExpectBegin()
	mock.ExpectExec("CREATE TABLE").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	database := &Database{conn: db}

	_, err = database.ExecuteSQLFromFile(filePath)
	if err != nil {
		t.Errorf("Error was not expected while executing SQL from file: %s", err)
	}

	if mockErr := mock.ExpectationsWereMet(); mockErr != nil {
		t.Errorf("Not all expected actions were completed: %s", err)
	}
}

func TestExecuteSQLFromFile_UpdateQuery(t *testing.T) {
	db, mock, err := setupMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	query := "UPDATE users SET name = 'Jane Doe' WHERE id = 1"
	filePath, err := createTempSQLFile(query)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filePath)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE users").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	database := &Database{conn: db}

	result, err := database.ExecuteSQLFromFile(filePath)
	if err != nil {
		t.Errorf("Error was not expected while executing SQL from file: %s", err)
	}

	rowsAffected, ok := result.(int64)
	if !ok || rowsAffected != 1 {
		t.Errorf("Expected 1 row to be affected, got %v", rowsAffected)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Not all expected actions were completed: %s", err)
	}
}

func TestExecuteSQLFromFile_DeleteQuery(t *testing.T) {
	db, mock, err := setupMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	query := "DELETE FROM users WHERE id = 1"
	filePath, err := createTempSQLFile(query)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filePath)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM users").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	database := &Database{conn: db}

	result, err := database.ExecuteSQLFromFile(filePath)
	if err != nil {
		t.Errorf("Error was not expected while executing SQL from file: %s", err)
	}

	rowsAffected, ok := result.(int64)
	if !ok || rowsAffected != 1 {
		t.Errorf("Expected 1 row to be affected, got %v", rowsAffected)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Not all expected actions were completed: %s", err)
	}
}

func TestExecuteSQLFromFile_UnknownQueryType(t *testing.T) {
	db, mock, err := setupMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	query := "SOME UNKNOWN QUERY"
	filePath, err := createTempSQLFile(query)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filePath)

	database := &Database{conn: db}

	_, err = database.ExecuteSQLFromFile(filePath)
	if err == nil {
		t.Errorf("Error was expected for unknown query type, but none occurred")
	}

	if errCheck := mock.ExpectationsWereMet(); errCheck != nil {
		t.Errorf("There were unexpected actions performed on the database: %s", err)
	}
}
