package cart

import "time"

type AddtoCartReq struct {
	UserID       int           `json:"user_id"`
	ProductItems []ProductItem `json:"product_items"`
}

type ProductItem struct {
	ID        int       `db:"id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	ProductID int       `db:"food_item_id" json:"product_id"`
	Quantity  int       `db:"quantity" json:"quantity"`
	UnitPrice float64   `db:"unit_price" json:"unit_price"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type AddToCartRes struct {
	Message string `json:"message"`
}

type DeleteFromCartReq struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
}

type DeleteFromCartRes struct {
	Message string `json:"message"`
}
