package main

import (
	"fmt"
	"net/http"

	"github.com/deepesh2508/go-cricket-web/env"
	"github.com/deepesh2508/go-cricket-web/handlers/cart"
	"github.com/deepesh2508/go-cricket-web/handlers/invoice"
	"github.com/deepesh2508/go-cricket-web/handlers/orders"
	"github.com/deepesh2508/go-cricket-web/handlers/users"
	m "github.com/deepesh2508/go-cricket-web/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	router := gin.New()
	router.RedirectFixedPath = true
	router.RedirectTrailingSlash = true

	// Health check route
	router.GET("/healthz", m.Healthz)

	// Apply middlewares
	router.Use(m.GenerateUUID(), m.RequestLogger(), gin.CustomRecovery(m.GinPanicRecovery()))

	// Define your routes and groups here

	// Cart routes
	cartGroup := router.Group("/cart")
	{
		cartGroup.POST("/add", cart.AddToCart)
		// Add other cart-related routes here
	}

	// Invoice routes
	invoiceGroup := router.Group("/invoice")
	{
		invoiceGroup.POST("/generate", invoice.GenerateInvoice)
		invoiceGroup.POST("/download", invoice.DownloadInvoice)
	}

	// Orders routes
	ordersGroup := router.Group("/orders")
	{
		ordersGroup.POST("/buynow", orders.BuyNow)
		ordersGroup.POST("/create-order", orders.CreateOrderFromCart)
	}

	// Users routes
	usersGroup := router.Group("/users")
	{
		usersGroup.POST("/signup", users.SignUp)
		usersGroup.POST("/signup", users.Login)
	}

	// Start the server
	server := &http.Server{
		Addr:    ":" + env.ENV.API_PORT,
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
