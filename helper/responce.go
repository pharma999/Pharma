package helper

import "github.com/gin-gonic/gin"

func SucessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"sucess": true,
		"data":   data,
	})
}

func ErrorResponce(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"sucess": false,
		"error":  message,
	})
}
