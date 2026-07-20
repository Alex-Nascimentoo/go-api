package main

import (
	"github.com/alex-nascimentoo/go-api/controller"
	"github.com/alex-nascimentoo/go-api/db"
	"github.com/alex-nascimentoo/go-api/repository"
	"github.com/alex-nascimentoo/go-api/routes"
	"github.com/alex-nascimentoo/go-api/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Connect to the database
	dbConnection, err := db.Connect()

	if err != nil {
		panic(err)
	}

	// Repositories
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Use cases
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// Controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Product routes
	routes.ProductRouter(server, ProductController)

	server.Run(":8000")
}
