package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/keyboards", getKeyboards)
	router.GET("/keyboards/:id", getKeyboardByID)
	router.POST("/keyboards", postKeyboards)

	router.Run(":3001")
}
