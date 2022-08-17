package services

import( 
	"context"
	"encoding"

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
	_, err := u.usercollection.Insert(u.ctx, user)
	return err
}

func (u *UserServiceimpl) GetUser(name *string)(*models.User, error){
	var user *models.User
	query := bson.D(bson.E([Kay: "name", Value: name]))
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	//return nil, nil
	return user, err
}

func (u *UserServiceimpl) UpdateUser(user *models.User) error{
	return nil
}

func (u *UserServiceimpl) DeleteUser(name *string) error{
	return nil
}