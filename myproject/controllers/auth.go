package controllers

import (
	"context"
	"log"
	"myproject/config"
	"myproject/models"
	"myproject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Register handles user registration
func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email is already registered
	existingUser := models.User{}
	filter := bson.M{"email": user.Email}
	err := config.UserCollection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Generate confirmation token
	user.ConfirmationToken = utils.GenerateToken()

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	// Insert user into database
	_, err = config.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Send confirmation email
	err = utils.SendConfirmationEmail(user.Email, user.ConfirmationToken)
	if err != nil {
		log.Printf("Failed to send confirmation email: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send confirmation email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully. Please check your email for confirmation."})
}
