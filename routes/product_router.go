package routes

import (
	"github.com/alex-nascimentoo/go-api/controller"
	"github.com/gin-gonic/gin"
)

func ProductRouter(server *gin.Engine, productController controller.ProductController) {
	server.GET("/products", productController.GetProducts)

	server.GET("/product/:id", productController.GetProductById)

	server.POST("/product", productController.CreateProduct)
}
