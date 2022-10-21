package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/segmentio/analytics-go"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"segment/leos-music-shop-api-go/data"
	"segment/leos-music-shop-api-go/docs"
	"segment/leos-music-shop-api-go/routes"
	"segment/leos-music-shop-api-go/segment"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	segment.SegmentClient = analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))
	data.Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func main() {
	data.Migrate()

	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	keyboardsGroup := router.Group("/keyboards")
	{
		keyboardsGroup.GET(":id", routes.GetKeyboardByID)
		keyboardsGroup.GET("", routes.GetKeyboards)
		keyboardsGroup.POST("", routes.PostKeyboard)
		// keyboardsGroup.DELETE(":id", c.DeleteAccount)
		// keyboardsGroup.PATCH(":id", c.UpdateAccount)
		// keyboardsGroup.POST(":id/images", c.UploadAccountImage)
	}

	manufacturersGroup := router.Group("/manufacturers")
	{
		manufacturersGroup.GET(":id", routes.GetManufacturerByID)
		manufacturersGroup.GET("", routes.GetManufacturers)
		manufacturersGroup.POST("", routes.PostManufacturer)
		// manufacturersGroup.DELETE(":id", c.DeleteAccount)
		// manufacturersGroup.PATCH(":id", c.UpdateAccount)
		// manufacturersGroup.POST(":id/images", c.UploadAccountImage)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	defer segment.SegmentClient.Close()
}
