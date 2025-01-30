package infraestructure

import (
	"demob/src/products/application"
	"demob/src/products/domain"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	createProductUseCase   *application.CreateProductUseCase
	viewAllProductsUseCase *application.ViewAllProductsUseCase
	updateProductUseCase   *application.UpdateProductUseCase
	deleteProductUseCase   *application.DeleteProductUseCase
}

func NewProductController(createUseCase *application.CreateProductUseCase, viewUseCase *application.ViewAllProductsUseCase, updateUseCase *application.UpdateProductUseCase, deleteUseCase *application.DeleteProductUseCase) *ProductController {
	return &ProductController{
		createProductUseCase:   createUseCase,
		viewAllProductsUseCase: viewUseCase,
		updateProductUseCase:   updateUseCase,
		deleteProductUseCase:   deleteUseCase,
	}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.createProductUseCase.Run(product.Nombre, product.Precio, product.Cantidad); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Producto creado correctamente",
		"product": gin.H{
			"nombre":   &product.Nombre,
			"cantidad": &product.Cantidad,
			"precio":   &product.Precio,
		},
	})
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	products, err := pc.viewAllProductsUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Id = int32(id)

	if err := pc.updateProductUseCase.Run(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := pc.deleteProductUseCase.Run(int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
