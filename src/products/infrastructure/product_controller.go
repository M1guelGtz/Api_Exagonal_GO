package infrastructure

import (
	"demob/src/products/application"
	"demob/src/products/infrastructure/broker"
	"demob/src/products/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	createProductUseCase   *handlers.CreateProductHandler
	viewAllProductsUseCase *handlers.GetAllProductsHandler
	updateProductUseCase   *handlers.UpdateProductHandler
	deleteProductUseCase   *handlers.DeleteProductHandler
	viewById               *handlers.GetProductByIdHandler
}

func NewProductController(
	createUseCase *application.CreateProductUseCase,
	viewUseCase *application.ViewAllProductsUseCase,
	updateUseCase *application.UpdateProductUseCase,
	deleteUseCase *application.DeleteProductUseCase,
	viewById *application.ViewProdByIdUseCase,
	publisher *broker.RabbitMQPublisher, // 🔹 Se agrega el publisher de RabbitMQ
) *ProductController {
	// 🔹 Se pasa el publisher al handler de creación
	createHandler := handlers.NewCreateProductHandler(createUseCase, publisher)
	viewHandler := handlers.NewGetAllProductsHandler(viewUseCase)
	updateHandler := handlers.NewUpdateProductHandler(updateUseCase)
	deleteHandler := handlers.NewDeleteProductHandler(deleteUseCase)
	viewByIdHandler := handlers.NewGetProductByIdHandler(viewById)

	return &ProductController{
		createProductUseCase:   createHandler,
		viewAllProductsUseCase: viewHandler,
		updateProductUseCase:   updateHandler,
		deleteProductUseCase:   deleteHandler,
		viewById:               viewByIdHandler,
	}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	pc.createProductUseCase.Handle(c)
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	pc.viewAllProductsUseCase.Handle(c)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	pc.updateProductUseCase.Handle(c)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	pc.deleteProductUseCase.Handle(c)
}

func (pc *ProductController) GetProductById(c *gin.Context) {
	pc.viewById.Handle(c)
}
