package main

import (
	"demob/src/products/application"
	"demob/src/products/infraestructure"
	"demob/src/users/application_users"
	domainusers "demob/src/users/domain_users"
	"demob/src/users/infraestructure_Users"
	"log"

	//"os"
	"github.com/gin-gonic/gin"
)

func main() {
	/*
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		name := os.Getenv("DB_NAME")
	*/
	dsn := ("root:passw0rd@tcp(127.0.0.1:8080)/Demo?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := infraestructure.NewMySql(dsn)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	dbus, err := infraestructureusers.NewMySQL(dsn)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	var userRepo domainusers.UserInterface = dbus 

	viewProductById := application.NewViewPrByIdUseCase(db)
	createProductUseCase := application.NewCreateProductUseCase(db)
	viewAllProductsUseCase := application.NewViewAllProductsUseCase(db)
	updateProductuseCase := application.NewUpdateProductUseCase(db)
	deleteProductUseCase := application.NewDeleteProductUseCase(db)
	createUserUseCase := applicationusers.NewCreateUserUseCase(userRepo)
	viewAllUsersUseCase := applicationusers.NewViewAllusersUseCase(userRepo)
	viewUserById := applicationusers.NewViewUserUseCase(userRepo)
	updateUserUseCase := applicationusers.NewUpdateUserUseCase(userRepo)
	deleteUserUseCase := applicationusers.NewDeleteUserUseCase(userRepo)
	userController := infraestructureusers.NewUserController(createUserUseCase, viewAllUsersUseCase, (*applicationusers.UpdateUserUseCase)(viewUserById),(*applicationusers.DeleteUserUseCase)(updateUserUseCase), (*applicationusers.ViewUserByIdUseCase)(deleteUserUseCase) )
	prodController := *infraestructure.NewProductController((*application.ViewProdByIdUseCase)(createProductUseCase), (*application.CreateProductUseCase)(viewAllProductsUseCase), (*application.ViewAllProductsUseCase)(updateProductuseCase), (*application.UpdateProductUseCase)(viewProductById), deleteProductUseCase)
	//productController := infraestructure.NewProductController(createProductUseCase, viewAllProductsUseCase, updateProductuseCase, deleteProductUseCase, viewProductById)
	r := gin.Default()
	infraestructure.RegisterProductRoutes(r, &prodController)
	infraestructureusers.RegisterUserRoutes(r, userController)
	if err := r.Run(":" + "3000"); err != nil {
		panic(err)
	}
}
