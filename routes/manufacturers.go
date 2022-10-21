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

// GetManufacturers godoc
// @Summary Gets all manufacturers
// @Schemes
// @Description Gets all manufacturers
// @Tags Manufacturers
// @Accept json
// @Produce json
// @Success 200
// @Router /manufacturers [get]
func GetManufacturers(c *gin.Context) {
	segment.SegmentClient.Enqueue(analytics.Track{
		UserId: "test-user",
		Event:  "All Manufacturers Listed",
	})
	data.Db.Find(&manufacturers)
	c.IndentedJSON(http.StatusOK, manufacturers)
}

// PostManufacturer godoc
// @Summary Creates a new Manufacturer
// @Schemes
// @Description Creates a new Manufacturer
// @Tags Manufacturers
// @Accept json
// @Produce json
// @Success 201
// @Router /manufacturers [post]
func PostManufacturer(c *gin.Context) {
	var newManufacturer models.Manufacturer

	if err := c.BindJSON(&newManufacturer); err != nil {
		return
	}

	manufacturers = append(manufacturers, newManufacturer)
	c.IndentedJSON(http.StatusCreated, newManufacturer)
}

// GetManufacturerByID godoc
// @Summary Gets a manufacturer by id
// @Schemes
// @Description Gets a manufacturer by id
// @Tags Manufacturers
// @Accept json
// @Produce json
// @Success 200
// @Router /manufacturers/{id} [get]
// @Param id path int true "Manufacturer id"
func GetManufacturerByID(c *gin.Context) {
	id := c.Param("id")

	var manufacturer models.Manufacturer
	result := data.Db.Find(&manufacturer, id)
	if result.RowsAffected > 0 {
		c.IndentedJSON(http.StatusOK, manufacturer)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Manufacturer not found"})
}
