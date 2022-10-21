package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/analytics-go"

	"segment/leos-music-shop-api-go/data"
	"segment/leos-music-shop-api-go/models"
	"segment/leos-music-shop-api-go/segment"
)

var manufacturers []models.Manufacturer

func GetManufacturers(c *gin.Context) {
	segment.SegmentClient.Enqueue(analytics.Track{
		UserId: "test-user",
		Event:  "All Manufacturers Listed",
	})
	data.Db.Find(&manufacturers)
	c.IndentedJSON(http.StatusOK, manufacturers)
}

func PostManufacturer(c *gin.Context) {
	var newManufacturer models.Manufacturer

	if err := c.BindJSON(&newManufacturer); err != nil {
		return
	}

	manufacturers = append(manufacturers, newManufacturer)
	c.IndentedJSON(http.StatusCreated, newManufacturer)
}

func GetManufacturerByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range manufacturers {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Manufacturer not found"})
}
