package main

import (
	"fmt"
	"net/http"

	"github.com/deepesh2508/go-cricket-web/env"
	"github.com/deepesh2508/go-cricket-web/handlers/cart"
	"github.com/deepesh2508/go-cricket-web/handlers/invoice"
	"github.com/deepesh2508/go-cricket-web/handlers/orders"
	"github.com/deepesh2508/go-cricket-web/handlers/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Cart routes
	cartGroup := router.Group("/cart")
	{
		cartGroup.POST("/add", cart.AddToCart)
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
		usersGroup.POST("/login", users.Login)
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
