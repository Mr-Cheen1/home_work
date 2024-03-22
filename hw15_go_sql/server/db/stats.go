package db

func GetUserStats(userID int) (map[string]interface{}, error) {
	var totalOrders int
	var totalAmount, avgProductPrice float64

	err := db.QueryRow("SELECT COUNT(*) FROM Orders WHERE user_id = $1", userID).Scan(&totalOrders)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow("SELECT SUM(total_amount) FROM Orders WHERE user_id = $1", userID).Scan(&totalAmount)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow(`
		 SELECT AVG(p.price) 
		 FROM Products p
		 JOIN OrderProducts op ON p.id = op.product_id
		 JOIN Orders o ON op.order_id = o.id
		 WHERE o.user_id = $1
	`, userID).Scan(&avgProductPrice)
	if err != nil {
		return nil, err
	}

	stats := map[string]interface{}{
		"total_orders":      totalOrders,
		"total_amount":      totalAmount,
		"avg_product_price": avgProductPrice,
	}
	return stats, nil
}
