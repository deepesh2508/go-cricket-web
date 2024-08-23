package cart

import (
	"github.com/deepesh2508/go-cricket-web/helper/errs"
)

var (
	ErrInvalidRequestPayload = errs.NewError(400, "IOS-CRT-USR-0001-E", "Invalid request JSON format")
	ErrNoProducts            = errs.NewError(400, "IOS-CRT-USR-0002-E", "No products found to add in cart")
	ErrProductNotFound       = errs.NewError(400, "IOS-CRT-USR-0003-E", "Product not found")
	ErrFetchingStock         = errs.NewError(500, "IOS-CRT-USR-0004-E", "Error fetching stock")
	ErrInsuffStock           = errs.NewError(400, "IOS-CRT-USR-0005-E", "Insufficient stock")
	ErrAddToCart             = errs.NewError(500, "IOS-CRT-USR-0006-E", "Error adding item to cart")
	ErrDeleteFromCart        = errs.NewError(500, "IOS-CRT-USR-0007-E", "Error in Deleting item From cart")
	ErrNoProductFound        = errs.NewError(500, "IOS-CRT-USR-0008-E", "Product not found in cart")
)
