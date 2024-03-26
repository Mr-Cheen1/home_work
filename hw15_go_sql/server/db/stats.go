package db

import (
	"database/sql"
	"errors"
)

type UserStats struct {
	TotalOrders     int     `json:"totalOrders"`
	TotalAmount     float64 `json:"totalAmount"`
	AvgProductPrice float64 `json:"avgProductPrice"`
}

func GetUserStats(userID *int) (*UserStats, error) {
	var stats UserStats

	tx, err := DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var query string
	var args []interface{}

	if userID != nil {
		query = `
            SELECT 
                COUNT(*) AS total_orders,
                COALESCE(SUM(o.total_amount), 0) AS total_amount,
                COALESCE(AVG(p.price), 0) AS avg_product_price
            FROM Orders o
            LEFT JOIN OrderProducts op ON o.id = op.order_id
            LEFT JOIN Products p ON op.product_id = p.id
            WHERE o.user_id = $1
        `
		args = []interface{}{*userID}
	} else {
		query = `
            SELECT 
                COUNT(*) AS total_orders,
                COALESCE(SUM(o.total_amount), 0) AS total_amount,
                COALESCE(AVG(p.price), 0) AS avg_product_price
            FROM Orders o
            LEFT JOIN OrderProducts op ON o.id = op.order_id
            LEFT JOIN Products p ON op.product_id = p.id
        `
	}

	err = tx.QueryRow(query, args...).Scan(
		&stats.TotalOrders,
		&stats.TotalAmount,
		&stats.AvgProductPrice,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &UserStats{}, nil
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &stats, nil
}
