package users

import (
	"log"
	"net/http"
	"time"

	"github.com/deepesh2508/go-cricket-web/database"
	"github.com/gin-gonic/gin"
)

// to get a new user onboard
func SignUp(c *gin.Context) {
	var req Signupreq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Mobile:    req.Mobile,
		Password:  hashedPassword,
		Role:      req.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `INSERT INTO users (firstname, lastname, email, password, mobile, role, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err = database.DB.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.Mobile, user.Role, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// to login a user
func Login(c *gin.Context) {
	var req Loginreq

	var user User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `SELECT * FROM users WHERE email = $1`
	err := database.DB.Get(&user, query, req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
	})
}
