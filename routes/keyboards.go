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
// @Summary get all keyboards
// @Schemes
// @Description get all keyboards
// @Tags example
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

func PostKeyboard(c *gin.Context) {
	var newKeyboard models.Keyboard

	if err := c.BindJSON(&newKeyboard); err != nil {
		return
	}

	keyboards = append(keyboards, newKeyboard)
	c.IndentedJSON(http.StatusCreated, newKeyboard)
}

func GetKeyboardByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range keyboards {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Keyboard not found"})
}
