package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/analytics-go"
)

var client analytics.Client

func main() {
	router := gin.Default()
	router.GET("/keyboards", getKeyboards)
	router.GET("/keyboards/:id", getKeyboardByID)
	router.POST("/keyboards", postKeyboards)

	router.Run(":3001")

	client := analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))
	defer client.Close()
}
