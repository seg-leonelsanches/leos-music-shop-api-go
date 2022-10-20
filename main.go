package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type keyboard struct {
	Id           string  `json:"id"`
	Model        string  `json:"model"`
	Manufacturer string  `json:"manufacturer"`
	Price        float64 `json:"price"`
}

var keyboards = []keyboard{
	{Id: "1", Model: "Williams Allegro III", Manufacturer: "Williams", Price: 349.99},
	{Id: "2", Model: "Yamaha P-125", Manufacturer: "Yamaha", Price: 699.99},
	{Id: "3", Model: "Casio CDP-S100", Manufacturer: "Casio", Price: 449.99},
}

func getKeyboards(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, keyboards)
}

func main() {
	router := gin.Default()
	router.GET("/keyboards", getKeyboards)

	router.Run(":3001")
}
