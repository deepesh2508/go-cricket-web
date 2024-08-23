package orders

import "github.com/deepesh2508/go-cricket-web/helper/errs"

var (
	ErrInvalidRequestPayload = errs.NewError(400, "IOS-ORD-ORD-0001-E", "Invalid request JSON format")
	ErrFetchingStock         = errs.NewError(500, "IOS-ORD-ORD-0002-E", "Error fetching stock")
	ErrScanCart              = errs.NewError(500, "IOS-ORD-ORD-0003-E", "Error In Scanning cart")
	ErrIteratingCart         = errs.NewError(500, "IOS-ORD-ORD-0004-E", "Error iterating over cart")
	ErrCreateOrder           = errs.NewError(500, "IOS-ORD-ORD-0005-E", "Error in creating order")
	ErrInsertOrderItems      = errs.NewError(500, "IOS-ORD-ORD-0006-E", "Error in inserting in order items")
	ErrDeleteFromCart        = errs.NewError(500, "IOS-ORD-ORD-0007-E", "Error in Deleting item From cart")
	ErrUpdateHistory         = errs.NewError(500, "IOS-ORD-ORD-0008-E", "Error in Updating purchase history")
	ErrFetchingProductPrice  = errs.NewError(500, "IOS-ORD-ORD-0009-E", "Error in Fetching product price")
)
