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
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message":err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message":err.Error()})
		return
	}
	//return ctx.JSON(200, "")
	return ctx.JSON(http.StatusOK,  gin.H{"message":"success"})
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

//Routes
func (uc *UserController) RegisterUserRoutes(rg gin.RouterGroup){
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/delete", uc.DeleteUser)
}