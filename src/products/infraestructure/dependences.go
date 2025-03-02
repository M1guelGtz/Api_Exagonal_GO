package infraestructure

import (
	"demob/src/products/application"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	ps := NewMySQL()

	createProduct := application.NewCreateProductUseCase(ps)
	getAllProducts := application.NewViewAllProductsUseCase(ps)
	updateProduct := application.NewUpdateProductUseCase(ps)
	deleteProduct := application.NewDeleteProductUseCase(ps)
	getById := application.NewViewPrByIdUseCase(ps)

	productController := NewProductController(createProduct, getAllProducts, updateProduct, deleteProduct, getById)
	
	RegisterProductRoutes(r, productController)
}
