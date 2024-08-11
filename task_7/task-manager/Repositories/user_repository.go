package repositories

import (
	"context"
	"errors"

	"github.com/kidistbezabih/task-manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = []byte("secretsecret")

type UserRepoManagement interface {
	RegisterUser(user *Domain.User) error
	Login(user *Domain.User) error
	Promote(user *Domain.User) error
	FindOneByUsername(user *Domain.User) error
	CountCollection() int64
	InsertOne(user Domain.User) error
	UpdateOneByUsername(username string) *mongo.UpdateResult
}

type UserServices struct {
	UsersCollection *mongo.Collection
	ctx             context.Context
}

func CreateNewUser(userscollection *mongo.Collection, ctx context.Context) UserServices {
	return UserServices{
		UsersCollection: userscollection,
		ctx:             ctx,
	}
}

func (us *UserServices) InsertOne(user Domain.User) error {
	_, err := us.UsersCollection.InsertOne(us.ctx, user)
	return err
}

func (us *UserServices) CountCollection() int64 {
	count, _ := us.UsersCollection.CountDocuments(context.Background(), bson.D{})
	return count
}

func (us *UserServices) FindOneByUsername(username string) error {
	var sample *Domain.User
	filter := bson.D{bson.E{Key: "username", Value: username}}
	err := us.UsersCollection.FindOne(us.ctx, filter).Decode(&sample)
	return err
}

func (us *UserServices) UpdateOneByUsername(username string) *mongo.UpdateResult {
	filter := bson.D{bson.E{Key: "username", Value: username}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "role", Value: "admin"}}}}
	result, _ := us.UsersCollection.UpdateOne(us.ctx, filter, update)
	return result

}
func (us *UserServices) RegisterUser(user *Domain.User) error {
	var sample *Domain.User
	filter := bson.D{bson.E{Key: "username", Value: user.UserName}}
	err := us.UsersCollection.FindOne(us.ctx, filter).Decode(&sample)
	if err == nil {
		return errors.New("user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	count, _ := us.UsersCollection.CountDocuments(context.Background(), bson.D{{}})
	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	_, err = us.UsersCollection.InsertOne(us.ctx, user)

	if err != nil {
		return err
	}
	return nil
}

func (us *UserServices) Login(user *Domain.User) (Domain.User, error) {

	var existingUser Domain.User
	filter := bson.D{bson.E{Key: "username", Value: user.UserName}}
	err := us.UsersCollection.FindOne(us.ctx, filter).Decode(&existingUser)
	if err != nil {
		return Domain.User{}, err
	}
	return existingUser, nil
}

func (us *UserServices) Promote(username string) error {
	filter := bson.D{bson.E{Key: "username", Value: username}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "role", Value: "admin"}}}}
	result, _ := us.UsersCollection.UpdateOne(us.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no user with this username")
	}
	return nil
}
