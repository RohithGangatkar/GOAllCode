package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//custom midddleware
	// router.Use(LoggerMiddleware())

	// Use our custom authentication middleware for a specific group of routes
	// apiRouter := router.Group("api")
	// apiRouter.Use(AuthMiddleware())

	//basic router
	router.GET("basic", func(ctx *gin.Context) {
		ctx.JSON(200, "OK")
	})

	// Route with URL parameters
	router.GET("/paramId/:id", getParamId())

	// Route with query parameters
	router.GET("/query/", getQeryId())

	//Route Groups
	// Public routes (no authentication required)
	public := router.Group("/public")
	{
		public.GET("/info", func(c *gin.Context) {
			c.String(200, "Public information")
		})
		public.GET("/products", func(c *gin.Context) {
			c.String(200, "Public product list")
		})
	}

	// Private routes (require authentication)
	private := router.Group("/private")
	private.Use(AuthMiddleware())
	{
		private.GET("/data", func(c *gin.Context) {
			c.String(200, "Private data accessible after authentication")
		})
		private.POST("/create", func(c *gin.Context) {
			c.String(200, "Create a new resource")
		})
	}

	//Controllers and Handlers
	//To improve code organization and maintainability,
	//Gin encourages using controllers to handle business logic separately from route handlers.
	userController := &UserController{}

	// Route using the UserController
	router.GET("/users/:id", userController.GetUserInfo)

	router.Run(":8080")

}

func getParamId() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"Id": id,
		})
	}
}

func getQeryId() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		c.String(200, "userId: "+id)
	}
}

type UserController struct{}

// GetUserInfo is a controller method to get user information
func (uc *UserController) GetUserInfo(c *gin.Context) {
	userID := c.Param("id")
	// Fetch user information from the database or other data source
	// For simplicity, we'll just return a JSON response.
	c.JSON(200, gin.H{"id": userID, "name": "John Doe", "email": "john@example.com"})
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	// In a real-world application, you would perform proper authentication here.
	// For the sake of this example, we'll just check if an API key is present.
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}
