package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/analytics-go"
)

type manufacturer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var manufacturers []manufacturer

func getManufacturers(c *gin.Context) {
	client.Enqueue(analytics.Track{
		UserId: "test-user",
		Event:  "All Manufacturers Listed",
	})
	c.IndentedJSON(http.StatusOK, manufacturers)
}

func postManufacturers(c *gin.Context) {
	var newManufacturer manufacturer

	if err := c.BindJSON(&newManufacturer); err != nil {
		return
	}

	manufacturers = append(manufacturers, newManufacturer)
	c.IndentedJSON(http.StatusCreated, newManufacturer)
}

func getManufacturerByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range manufacturers {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "manufacturer not found"})
}
