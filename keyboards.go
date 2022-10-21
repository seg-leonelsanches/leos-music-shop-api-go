package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/analytics-go"
)

type keyboard struct {
	Id           string  `json:"id"`
	Model        string  `json:"model"`
	Manufacturer string  `json:"manufacturer"`
	Price        float64 `json:"price"`
}

var keyboards []keyboard

func getKeyboards(c *gin.Context) {
	client.Enqueue(analytics.Track{
		UserId: "test-user",
		Event:  "All Keyboards Listed",
	})
	c.IndentedJSON(http.StatusOK, db.Find(&keyboards))
}

func postKeyboards(c *gin.Context) {
	var newKeyboard keyboard

	if err := c.BindJSON(&newKeyboard); err != nil {
		return
	}

	keyboards = append(keyboards, newKeyboard)
	c.IndentedJSON(http.StatusCreated, newKeyboard)
}

func getKeyboardByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range keyboards {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "keyboard not found"})
}
