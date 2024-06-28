package controller

import (
	"github.com/duonglc27092000/go-ecommerce-backend-api/internal/service"
	"github.com/duonglc27092000/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc -  user controller

// controller -> service -> repo -> models -> dbs
func (uc *UserController) GetUserByID(c *gin.Context) {

	// if err != nil {
	// 	return response.ErrorResponse(c, 20003, "No Need !")

	// }
	response.SuccessResponse(c, 20001, []string{"d1", "d22", "d333"})
	// response.SuccessResponse(c, 20001, []string{"d1", "d22", "d333"})

	// response.ErrorResponse(c, 20003, "No Need !")
}
