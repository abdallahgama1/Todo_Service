package handlers

import (
	"Todo_Service/models"
	"Todo_Service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Register godoc
// @Summary      Register a new user
// @Description  Create a new user account with a username and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input   body      models.RegisterInput  true  "Registration Info"
// @Success      200     {object}  map[string]string
// @Failure      400     {object}  map[string]string "Invalid input or user already exists"
// @Router       /auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.service.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

// Login godoc
// @Summary      Log in a user
// @Description  Log in with username and password to receive a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input   body      models.LoginInput  true  "Login Credentials"
// @Success      200     {object}  map[string]string
// @Failure      401     {object}  map[string]string "Invalid credentials"
// @Router       /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
