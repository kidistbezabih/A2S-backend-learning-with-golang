package usecases

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/kidistbezabih/task-manager/Domain"
	repositories "github.com/kidistbezabih/task-manager/Repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userservices repositories.UserServices
}

func NewUseCase(userservices repositories.UserServices) UserUseCase {
	return UserUseCase{
		userservices: userservices,
	}
}

func (us *UserUseCase) RegisterUser(user *Domain.User) error {
	err := us.userservices.FindOneByUsername(user.UserName)

	if err == nil {
		return errors.New("user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	count := us.userservices.CountCollection()
	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	err = us.userservices.InsertOne(*user)

	if err != nil {
		return err
	}
	return nil
}

func (us *UserUseCase) Login(user *Domain.User) (string, error) {
	// var existingUser Domain.User
	existingUser, err := us.userservices.Login(user)
	if err != nil {
		return "", err
	}
	err = us.userservices.FindOneByUsername(user.UserName)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		return "", errors.New("invalied username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": existingUser.UserName,
		"role":     existingUser.Role,
	})

	jwtToken, err := token.SignedString(repositories.SecretKey)

	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (us *UserUseCase) Promote(username string) error {
	result := us.userservices.UpdateOneByUsername(username)
	if result.MatchedCount != 1 {
		return errors.New("no user with this username")
	}
	return nil
}
