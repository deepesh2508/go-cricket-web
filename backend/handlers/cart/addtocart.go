package cart

import (
	"database/sql"
	"time"

	"github.com/deepesh2508/go-cricket-web/database"
	"github.com/deepesh2508/go-cricket-web/helper/util"
	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var req AddtoCartReq

	if err := c.ShouldBindJSON(&req); err != nil {
		util.SendResponse(c, nil, ErrInvalidRequestPayload, err)
		return
	}

	if len(req.ProductItems) <= 0 {
		util.SendResponse(c, nil, ErrNoProducts, nil)
		return
	}

	for _, item := range req.ProductItems {
		if item.Quantity <= 0 {
			util.SendResponse(c, nil, ErrNoProducts, nil)
			return
		}

		var stock int
		err := database.DB.QueryRow("SELECT stock FROM Products WHERE product_id = $1", item.ProductID).Scan(&stock)
		if err == sql.ErrNoRows {
			util.SendResponse(c, nil, ErrProductNotFound, nil)
			return
		} else if err != nil {
			util.SendResponse(c, nil, ErrFetchingStock, err)
			return
		}

		if stock < item.Quantity {
			util.SendResponse(c, nil, ErrInsuffStock, nil)
			return
		}

		cartItem := ProductItem{
			UserID:    req.UserID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UnitPrice: item.UnitPrice,
		}

		query := `INSERT INTO CartItems (user_id, product_id, quantity, created_at, updated_at, unit_price)
				  VALUES ($1, $2, $3, $4, $5, $6)
				  ON CONFLICT (user_id, product_id)
				  DO UPDATE SET quantity = CartItems.quantity + EXCLUDED.quantity, updated_at = EXCLUDED.updated_at
				  RETURNING id`
		err = database.DB.QueryRow(query, cartItem.UserID, cartItem.ProductID, cartItem.Quantity, cartItem.CreatedAt, cartItem.UpdatedAt, cartItem.UnitPrice).Scan(&cartItem.ID)
		if err != nil {
			util.SendResponse(c, nil, ErrAddToCart, err)
			return
		}

	}

	util.SendResponse(c, AddToCartRes{
		Message: "Successfully added to cart"},
		nil, nil)
}

func DeleteFromCart(c *gin.Context) {
	var req DeleteFromCartReq

	if err := c.ShouldBindJSON(&req); err != nil {
		util.SendResponse(c, nil, ErrInvalidRequestPayload, err)
		return
	}

	if req.ProductID <= 0 {
		util.SendResponse(c, nil, ErrNoProducts, nil)
		return
	}

	query := `DELETE FROM CartItems WHERE user_id = $1 AND product_id = $2`
	result, err := database.DB.Exec(query, req.UserID, req.ProductID)
	if err != nil {
		util.SendResponse(c, nil, ErrDeleteFromCart, err)
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {

	}

	if rowsAffected == 0 {
		util.SendResponse(c, nil, ErrNoProductFound, err)
		return
	}

	util.SendResponse(c, DeleteFromCartRes{
		Message: "Item removed from the cart"}, nil, nil)
}
