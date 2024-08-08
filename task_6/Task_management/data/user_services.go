package data

import (
	"context"
	"errors"
	"task_management/models"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = []byte("secretsecret")

type UserManagement interface {
	RegisterUser(user *models.User) error
	Login(user *models.User) error
	Promote(user *models.User) error
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

func (us *UserServices) RegisterUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "username", Value: us.UsersCollection}}
	exist := us.UsersCollection.FindOne(us.ctx, filter)
	if exist == nil {
		return errors.New("user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// count the collection
	// if 0
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

func (us *UserServices) Login(user *models.User) (string, error) {

	var existingUser models.User
	filter := bson.D{bson.E{Key: "username", Value: user.UserName}}
	err := us.UsersCollection.FindOne(us.ctx, filter).Decode(&existingUser)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		return "", errors.New("invalied username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": existingUser.UserName,
		"role":     existingUser.Role,
	})

	jwtToken, err := token.SignedString(SecretKey)

	if err != nil {
		return "", err
	}
	return jwtToken, nil
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

// func (us *UserServices) GetAUserTasks(username string) ([]models.Task, error) {
// 	filter := bson.D{bson.E{Key: "username", Value: username}}
// 	cursor, err := us.UsersCollection.Find(us.ctx, filter)

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(us.ctx)

// 	for cursor.Next(us.ctx) {
// 		var user models.Task

// 		err := cursor.Decode(&user)
// 		if err != nil {
// 			return []models.Task{}, err
// 		}

// 		tasks = append(tas)

// 	}

// }
