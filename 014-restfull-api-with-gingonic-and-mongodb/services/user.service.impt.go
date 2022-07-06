package services

import( 
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/FuaAlfu/Go-Projects/014-restfull-api-with-gingonic-and-mongodb/models"
)

type UserServiceimpl struct{
	usercollection *mongo.Collection
	ctx             context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService{
	return &UserServiceimpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceimpl) GetAll([] *models.User, error){
	return nil, nil
}

func (u *UserServiceimpl) CreateUser(user *models.User) error{
	return nil
}

func (u *UserServiceimpl) GetUser(name string)(*models.User, error){
	return nil, nil
}

func (u *UserServiceimpl) UpdateUser(user *models.User) error{
	return nil
}

func (u *UserServiceimpl) DeleteUser(name string) error{
	return nil
}