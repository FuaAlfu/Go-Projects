package controller

import (
	"github.com/FuaAlfu/Go-Projects/014-restfull-api-with-gingonic-and-mongodb/services"
	"github.com/gin-gonic/gin"

)

func New(userservice services.UserService) UserController {
	//constructing..
	return UserController{
		UserService: userservice,
	}
}

func (u *UserServiceimpl) GetAll() {
	return nil, nil
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	return ctx.JSON(200, "")
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	return ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	return ctx.JSON(200, "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	return ctx.JSON(200, "")
}
