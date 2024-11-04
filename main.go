package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "log"
    "wav-to-flac-conversion/controllers"
)

func main() {
    // Initialize Gin router
    router := gin.Default()

    // Configure CORS middleware
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8080"}, // Adjust to your frontend origin
        AllowMethods:     []string{"GET", "POST"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        AllowCredentials: true,
    }))

    // Set up HTTP and WebSocket routes
    router.POST("/upload", controllers.HandleWavUpload)
    router.GET("/ws", func(c *gin.Context) {
        controllers.HandleWebSocket(c.Writer, c.Request)
    })

    router.GET("/", func(c *gin.Context){
        c.File("./static/index.html")
    })

    // Serve static files for frontend
    router.Static("/static", "./static")
	router.Static("/download", "./temp")

    // Start server on port 8080
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}