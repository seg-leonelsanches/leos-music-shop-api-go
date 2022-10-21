package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/analytics-go"

	"segment/leos-music-shop-api-go/data"
	"segment/leos-music-shop-api-go/models"
	"segment/leos-music-shop-api-go/segment"
)

var keyboards []models.Keyboard

// GetKeyboards godoc
// @Summary Gets all keyboards
// @Schemes
// @Description Gets all keyboards
// @Tags Keyboards
// @Accept json
// @Produce json
// @Success 200
// @Router /keyboards [get]
func GetKeyboards(c *gin.Context) {
	segment.SegmentClient.Enqueue(analytics.Track{
		UserId: "test-user",
		Event:  "All Keyboards Listed",
	})
	data.Db.Find(&keyboards)
	c.IndentedJSON(http.StatusOK, keyboards)
}

// PostKeyboard godoc
// @Summary Creates a new Keyboard
// @Schemes
// @Description Creates a new Keyboard
// @Tags Keyboards
// @Accept json
// @Produce json
// @Success 201
// @Router /keyboards [post]
func PostKeyboard(c *gin.Context) {
	var newKeyboard models.Keyboard

	if err := c.BindJSON(&newKeyboard); err != nil {
		return
	}

	keyboards = append(keyboards, newKeyboard)
	c.IndentedJSON(http.StatusCreated, newKeyboard)
}

// GetKeyboardByID godoc
// @Summary Gets a keyboard by id
// @Schemes
// @Description Gets a keyboard by id
// @Tags Keyboards
// @Accept json
// @Produce json
// @Success 200
// @Router /keyboards/{id} [get]
// @Param id path int true "Keyboard id"
func GetKeyboardByID(c *gin.Context) {
	id := c.Param("id")

	var keyboard models.Keyboard
	result := data.Db.Find(&keyboard, id)
	if result.RowsAffected > 0 {
		c.IndentedJSON(http.StatusOK, keyboard)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Keyboard not found"})
}
