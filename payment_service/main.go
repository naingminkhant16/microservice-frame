package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/payment", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello From Payment service"})
	})

	router.GET("/api/payment/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"payment": "User payment",
		})
	})

	router.GET("/api/payment/user/details", func(ctx *gin.Context) {
		// Call user service to get user details
		resp, err := http.Get("http://user_service:3000/api/user")
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to fetch user details"})
			return
		}
		defer resp.Body.Close()

		// For simplicity, just return a message; in real code, parse the response
		ctx.JSON(200, gin.H{
			"payment":               "User payment details",
			"user_service_response": "Fetched user data",
		})
	})

	router.Run()
}
