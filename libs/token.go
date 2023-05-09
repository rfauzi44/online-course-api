package libs

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var Secret = []byte(os.Getenv("JWT_KEY"))

func GenerateToken(id, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CheckToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
