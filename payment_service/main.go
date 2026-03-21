package main

import "github.com/gin-gonic/gin"

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

	router.Run()
}
