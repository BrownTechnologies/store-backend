package main

import (
	"store-backend/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", controller.Health)
	router.PUT("/update-zip-codes", controller.UpdateZipCodes)

	router.Run(":8000")

}
