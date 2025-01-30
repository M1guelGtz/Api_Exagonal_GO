package infraestructureusers

import (
	applicationusers "demob/src/users/application_users"
	domainusers "demob/src/users/domain_users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	createUserUsecase   *applicationusers.CreateUserUseCase
	viewUserUseCase     *applicationusers.ViewUserByIdUseCase
	viewAllUsersUseCase *applicationusers.ViewAllUsersUseCase
	updateUserUseCase   *applicationusers.UpdateUserUseCase
	deleteUserUseCase   *applicationusers.DeleteUserUseCase
}

func NewUserController(createUseCase *applicationusers.CreateUserUseCase, viewAllUseCase *applicationusers.ViewAllUsersUseCase, updateUseCase *applicationusers.UpdateUserUseCase, deleteUseCase *applicationusers.DeleteUserUseCase, viewUseCase *applicationusers.ViewUserByIdUseCase) *UserController {
	return &UserController{
		createUserUsecase:   createUseCase,
		viewAllUsersUseCase: viewAllUseCase,
		updateUserUseCase:   updateUseCase,
		deleteUserUseCase:   deleteUseCase,
		viewUserUseCase:     viewUseCase,
	}
}

func (pc *UserController) CreateUser(c *gin.Context) {
	var user domainusers.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.createUserUsecase.Execute(user.Name, user.Email, user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Usuario creado correctamente",
		"Usuario": gin.H{
			"name":     &user.Name,
			"email":    &user.Email,
			"password": &user.Password,
		},
	})
}

func (pc *UserController) GetAllUsers(c *gin.Context) {
	users, err := pc.viewAllUsersUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
func (pc *UserController) GetUserById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := pc.viewUserUseCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
func (pc *UserController) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user domainusers.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = int32(id)

	if err := pc.updateUserUseCase.Execute(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func (pc *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := pc.deleteUserUseCase.Execute(int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
