package infrastructure

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	repositories "github.com/kidistbezabih/task-manager/Repositories"
)

func validateSigningMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return repositories.SecretKey, nil
}
