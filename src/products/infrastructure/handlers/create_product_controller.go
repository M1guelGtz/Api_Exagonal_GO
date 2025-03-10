package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"demob/src/products/application"
	"demob/src/products/domain"
	"demob/src/products/infrastructure/broker"

	"github.com/gin-gonic/gin"
)

type CreateProductHandler struct {
	UseCase   *application.CreateProductUseCase
	Publisher *broker.RabbitMQPublisher
}

func NewCreateProductHandler(useCase *application.CreateProductUseCase, publisher *broker.RabbitMQPublisher) *CreateProductHandler {
	return &CreateProductHandler{UseCase: useCase, Publisher: publisher}
}

func (h *CreateProductHandler) Handle(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UseCase.Run(product.Nombre, product.Precio, product.Cantidad); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	message := map[string]interface{}{
		"nombre":   product.Nombre,
		"cantidad": product.Cantidad,
		"precio":   product.Precio,
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Println("Error al serializar el mensaje:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el mensaje"})
		return
	}
	err = h.Publisher.Publish(messageBytes)
	if err != nil {
		log.Println("Error publicando mensaje en RabbitMQ:", err)
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
