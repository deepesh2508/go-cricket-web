package orders

import (
	"time"

	"github.com/deepesh2508/go-cricket-web/database"
	"github.com/deepesh2508/go-cricket-web/helper/util"
	"github.com/gin-gonic/gin"
)

func CreateOrderFromCart(c *gin.Context) {
	var req CreateOrderReq

	if err := c.ShouldBindJSON(&req); err != nil {
		util.SendResponse(c, nil, ErrInvalidRequestPayload, err)
		return
	}

	var totalAmount float64
	rows, err := database.DB.Query(`
        SELECT ci.product_id, ci.quantity, p.price
        FROM CartItems ci
        JOIN Products p ON ci.product_id = p.product_id                 b
        WHERE ci.user_id = $1`, req.UserID)
	if err != nil {
		util.SendResponse(c, nil, ErrFetchingStock, err)
		return
	}
	defer rows.Close()

	var orderItems []OrderItem
	for rows.Next() {
		var item OrderItem
		var unitPrice float64
		if err := rows.Scan(&item.ProductID, &item.Quantity, &unitPrice); err != nil {
			util.SendResponse(c, nil, ErrScanCart, err)
			return
		}
		item.Price = unitPrice * float64(item.Quantity)
		orderItems = append(orderItems, item)
		totalAmount += item.Price
	}

	if err := rows.Err(); err != nil {
		util.SendResponse(c, nil, ErrIteratingCart, err)
		return
	}

	var orderID int

	//inserting into orders
	err = database.DB.QueryRow(`
        INSERT INTO Orders (user_id, total_amount, payment_method_id, shipping_address, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, 'pending', $5, $5) RETURNING order_id`,
		req.UserID, totalAmount, req.PaymentMethodID, req.ShippingAddress, time.Now()).Scan(&orderID)
	if err != nil {
		util.SendResponse(c, nil, ErrCreateOrder, err)
		return
	}

	//inserting into order items
	for _, item := range orderItems {
		_, err = database.DB.Exec(`
            INSERT INTO OrderItems (order_id, product_id, quantity, price, created_at)
            VALUES ($1, $2, $3, $4, $5)`,
			orderID, item.ProductID, item.Quantity, item.Price, time.Now())
		if err != nil {
			util.SendResponse(c, nil, ErrInsertOrderItems, err)
			return
		}
	}

	//deleting from cart items
	_, err = database.DB.Exec("DELETE FROM CartItems WHERE user_id = $1", req.UserID)
	if err != nil {
		util.SendResponse(c, nil, ErrDeleteFromCart, err)
		return
	}

	//Update purchase history in users table
	_, err = database.DB.Exec("UPDATE users SET purchase_history = purchase_history || $1 WHERE user_id = $2", "New purchase data", req.UserID)
	if err != nil {
		util.SendResponse(c, nil, ErrUpdateHistory, err)
		return
	}

	util.SendResponse(c, CreateOrderRes{
		Message: "Order Placed Successfully",
	}, nil, nil)
}

func BuyNow(c *gin.Context) {
	var req BuyNowReq

	if err := c.ShouldBindJSON(&req); err != nil {
		util.SendResponse(c, nil, ErrInvalidRequestPayload, err)
		return
	}

	var totalAmount float64
	var productPrice float64

	err := database.DB.QueryRow("SELECT price FROM products WHERE product_id = $1", req.ProductID).Scan(&productPrice)
	if err != nil {
		util.SendResponse(c, nil, ErrFetchingProductPrice, err)
		return
	}

	totalAmount = productPrice * float64(req.Quantity)
	orderItems := []OrderItem{
		{
			ProductID: req.ProductID,
			Quantity:  req.Quantity,
			Price:     productPrice,
		},
	}

	var orderID int

	err = database.DB.QueryRow(`
        INSERT INTO Orders (user_id, total_amount, payment_method_id, shipping_address, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, 'pending', $5, $5) RETURNING order_id`,
		req.UserID, totalAmount, req.PaymentMethodID, req.ShippingAddress, time.Now()).Scan(&orderID)
	if err != nil {
		util.SendResponse(c, nil, ErrCreateOrder, err)
		return
	}

	// Inserting into order items
	for _, item := range orderItems {
		_, err = database.DB.Exec(`
            INSERT INTO OrderItems (order_id, product_id, quantity, price, created_at)
            VALUES ($1, $2, $3, $4, $5)`,
			orderID, item.ProductID, item.Quantity, item.Price, time.Now())
		if err != nil {
			util.SendResponse(c, nil, ErrInsertOrderItems, err)
			return
		}
	}

	_, err = database.DB.Exec("UPDATE users SET purchase_history = purchase_history || $1 WHERE user_id = $2", "New purchase data", req.UserID)
	if err != nil {
		util.SendResponse(c, nil, ErrUpdateHistory, err)
		return
	}

	util.SendResponse(c, BuyNowRes{
		Message: "Order Placed Successfully",
	}, nil, nil)

}
