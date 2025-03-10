package infrastructure

import (
	"demob/src/products/application"
	"demob/src/products/infrastructure/broker"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func Init(r *gin.Engine) {
	ps := NewMySQL()

	
	err := godotenv.Load()
	createProduct := application.NewCreateProductUseCase(ps)
	getAllProducts := application.NewViewAllProductsUseCase(ps)
	updateProduct := application.NewUpdateProductUseCase(ps)
	deleteProduct := application.NewDeleteProductUseCase(ps)
	getById := application.NewViewPrByIdUseCase(ps)

	host := os.Getenv("BROKER_HOST")
	user := os.Getenv("BROKER_USER")
	pass := os.Getenv("BROKER_PASS")
	// Configurar RabbitMQ
	connURL := fmt.Sprintf("amqp://%s:%s@%s/", user, pass, host)
	conn, err := amqp.Dial(connURL)
	//conn, err := amqp.Dial("amqp://miguel:7s0725FLU2@50.19.162.154:5672/")
	if err != nil {
		log.Fatal("Error al conectar con RabbitMQ", err)
	}

	publisher, err := broker.NewRabbitMQPublisher(conn, "Q1")
	if err != nil {
		log.Fatal("Error al crear el publicador de RabbitMQ", err)
	}

	productController := NewProductController(createProduct, getAllProducts, updateProduct, deleteProduct, getById, publisher)

	RegisterProductRoutes(r, productController)
}
