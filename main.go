package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/segmentio/analytics-go"
)

var client analytics.Client

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client = analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))
}

func main() {
	router := gin.Default()
	router.GET("/keyboards", getKeyboards)
	router.GET("/keyboards/:id", getKeyboardByID)
	router.POST("/keyboards", postKeyboards)

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
