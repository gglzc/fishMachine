package controllers

import (
	"net/http"
	"strconv"

	"github.com/gglzc/fishMachine/service/user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func (uc *UserController) GetUser(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user, err := uc.UserService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}



