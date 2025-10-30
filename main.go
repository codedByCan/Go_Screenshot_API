package main

import (
	"log"
	"time"

	"github.com/codedByCan/Go_Screenshot_API/controllers/api"
	"github.com/codedByCan/Go_Screenshot_API/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())

	r.Use(middleware.CORSMiddleware())

	v1 := r.Group("/api/v1")
	{
		v1.POST("/screenshot", api.HandleScreenshot)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	log.Println("🚀 Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server başlatılamadı:", err)
	}
}
