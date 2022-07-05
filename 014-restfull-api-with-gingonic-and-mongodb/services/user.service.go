package services

import(
	 "github.com/FuaAlfu/Go-Projects/014-restfull-api-with-gingonic-and-mongodb/models"
)

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string)(*models.User, error)
	GetAll()([]*models.User, error)
	UpdateUser(*models.User)
	DeleteUser(*string) error
}