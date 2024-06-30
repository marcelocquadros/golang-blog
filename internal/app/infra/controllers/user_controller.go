package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcelocquadros/blog/internal/app/usecases/user"
)

type UserController struct {
	createUser   user.CreateUser
	deleteUser   user.DeleteUser
	findAllUsers user.FindAllUsers
	findUserByID user.FindUserByID
	updateUser   user.UpdateUser
}

func NewUserController(
	createUser user.CreateUser,
	deleteUser user.DeleteUser,
	findAll user.FindAllUsers,
	findUserByID user.FindUserByID,
	updateUser user.UpdateUser) *UserController {
	return &UserController{
		createUser:   createUser,
		deleteUser:   deleteUser,
		findAllUsers: findAll,
		findUserByID: findUserByID,
		updateUser:   updateUser,
	}
}

func (c UserController) CreateUser(ctx *gin.Context) *user.CreateUserResponse {
	return &user.CreateUserResponse{}
}

func (c UserController) DeleteUser(ctx *gin.Context) {
}

func (c UserController) FindAllUsers(ctx *gin.Context) *user.FindAllUsersResponse {
	return nil
}

func (c UserController) FindUserByID(ctx *gin.Context) *user.FindUserByIDResponse {
	return nil
}

func (c UserController) UpdateUser(ctx *gin.Context) *user.UpdateUserResponse {
	return nil
}
