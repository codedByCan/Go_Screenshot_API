package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"speedapi/controllers/api"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiGroup := r.Group("/api/v1")
	{
		apiGroup.POST("/screenshot", api.HandleScreenshot)
	}
	return r
}

func main() {
	router := setupRouter()
	if err := router.Run(":3000"); err != nil {
		log.Fatalf("Sunucu başlatılamadı: %v", err)
	}
}