package integration

import (
	"context"
	"myproject/config"
	"myproject/controllers"
	"myproject/models"
	"testing"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	config.InitDB()
	router.POST("/register", controllers.Register)
	router.GET("/confirm", controllers.ConfirmEmail)
	router.POST("/login", controllers.Login)
	router.GET("/users", controllers.GetAllUsers)
	return router
}

func TestRegisterIntegration(t *testing.T) {
	router := setupRouter()

	user := models.User{
		Nickname: "testuser",
		Email:    "testuser@example.com",
		Password: "password",
	}
	jsonUser, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User registered successfully")

	// Clean up
	filter := bson.M{"email": user.Email}
	_, err := config.UserCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		t.Fatalf("Failed to clean up test user: %v", err)
	}
}
