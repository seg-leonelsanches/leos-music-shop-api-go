package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/segmentio/analytics-go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var client analytics.Client
var db *gorm.DB

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client = analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func main() {

	// Migrate the schema
	db.AutoMigrate(&keyboard{})
	db.AutoMigrate(&manufacturer{})

	// Create
	db.Create([]keyboard{
		{Id: "1", Model: "Williams Allegro III", Manufacturer: "Williams", Price: 349.99},
		{Id: "2", Model: "Yamaha P-125", Manufacturer: "Yamaha", Price: 699.99},
		{Id: "3", Model: "Casio CDP-S100", Manufacturer: "Casio", Price: 449.99},
	})
	db.Create([]manufacturer{
		{Id: "1", Name: "Williams"},
		{Id: "2", Name: "Yamaha"},
		{Id: "3", Name: "Casio"},
	})

	router := gin.Default()
	router.GET("/keyboards", getKeyboards)
	router.GET("/keyboards/:id", getKeyboardByID)
	router.POST("/keyboards", postKeyboards)

	router.GET("/manufacturers", getManufacturers)
	router.GET("/manufacturers/:id", getManufacturerByID)
	router.POST("/manufacturers", postManufacturers)

	server := &http.Server{
		Addr:    ":3001",
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}

	log.Println("Server exiting")
	defer client.Close()
}
