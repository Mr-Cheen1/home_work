package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const queryTypeSelect = "SELECT"

type Database struct {
	conn *sql.DB
}

func NewDatabase(connString string) (*Database, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &Database{conn: db}, nil
}

func cleanAndDetectType(query string) (cleanQuery string, queryType string) {
	scanner := bufio.NewScanner(strings.NewReader(query))
	var result strings.Builder

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if idx := strings.Index(line, "--"); idx != -1 {
			line = strings.TrimSpace(line[:idx])
		}
		if len(line) == 0 {
			continue
		}
		result.WriteString(line + " ")
	}

	cleanedQuery := result.String()
	upperQuery := strings.ToUpper(cleanedQuery)
	switch {
	case strings.HasPrefix(upperQuery, queryTypeSelect):
		queryType = queryTypeSelect
	case strings.HasPrefix(upperQuery, "INSERT"):
		queryType = "INSERT"
	case strings.HasPrefix(upperQuery, "UPDATE"):
		queryType = "UPDATE"
	case strings.HasPrefix(upperQuery, "DELETE"):
		queryType = "DELETE"
	case strings.HasPrefix(upperQuery, "CREATE"):
		queryType = "CREATE"
	case strings.HasPrefix(upperQuery, "CREATE INDEX"):
		queryType = "INDEX"
	default:
		queryType = "UNKNOWN"
	}
	return cleanedQuery, queryType
}

func (db *Database) ExecuteSQLFromFile(filePath string) (interface{}, error) {
	// Чтение и очистка запроса из файла.
	cleanedQuery, queryType, err := db.readAndCleanQueryFromFile(filePath)
	if err != nil {
		return nil, err
	}

	tx, err := db.conn.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	switch queryType {
	case queryTypeSelect:
		return db.executeSelectQuery(tx, cleanedQuery)
	case "INSERT", "UPDATE", "DELETE":
		return db.executeModificationQuery(tx, cleanedQuery)
	case "CREATE", "INDEX":
		return nil, db.executeCreateQuery(tx, cleanedQuery)
	default:
		return nil, fmt.Errorf("unsupported query type: %s", queryType)
	}
}

func (db *Database) executeSelectQuery(tx *sql.Tx, query string) ([]map[string]interface{}, error) {
	rows, err := tx.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}
	for rows.Next() {
		columnsValues := make([]interface{}, len(columns))
		columnsPointers := make([]interface{}, len(columns))
		for i := range columnsValues {
			columnsPointers[i] = &columnsValues[i]
		}

		err = rows.Scan(columnsPointers...)
		if err != nil {
			return nil, err
		}

		result := make(map[string]interface{})
		for i, colName := range columns {
			val := columnsValues[i]
			b, ok := val.([]byte)
			if ok {
				result[colName] = string(b)
			} else {
				result[colName] = val
			}
		}
		results = append(results, result)
	}
	return results, nil
}

func (db *Database) executeModificationQuery(tx *sql.Tx, query string) (int64, error) {
	result, err := tx.Exec(query)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (db *Database) executeCreateQuery(tx *sql.Tx, query string) error {
	_, err := tx.Exec(query)
	return err
}

func (db *Database) readAndCleanQueryFromFile(filePath string) (string, string, error) {
	queryBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", "", err
	}
	query := string(queryBytes)

	cleanedQuery, queryType := cleanAndDetectType(query)
	return cleanedQuery, queryType, nil
}
