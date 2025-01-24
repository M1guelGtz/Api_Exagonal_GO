package infraestructure

import (
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, productController *ProductController) {
	products := r.Group("/products")
	{
		products.POST("/create", productController.CreateProduct)
		products.GET("/", productController.GetAllProducts)
		products.PUT("/update/:id", productController.UpdateProduct)
		products.DELETE("/delete/:id", productController.DeleteProduct)
	}
}
