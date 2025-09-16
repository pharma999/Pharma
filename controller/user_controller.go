package controller

import (
	"demo/database"
	"demo/helper"
	"demo/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		helper.ErrorResponce(c, "Invalid Input")
		return
	}

	if result := database.DB.Create(&user); result.Error != nil {
		helper.ErrorResponce(c, result.Error.Error())
		return
	}

	helper.SucessResponse(c, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	if result := database.DB.Find(&users); result.Error != nil {
		helper.ErrorResponce(c, result.Error.Error())
		return
	}
	helper.SucessResponse(c, users)
}
