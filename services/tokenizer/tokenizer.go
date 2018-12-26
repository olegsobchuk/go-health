package tokenizer

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secretKey = "some&super#sekret*key"

// BuildNew build new signed token
func BuildNew(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"nbf":    time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

// Parse prses encripted token to data
func Parse(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("can't parse token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return map[string]interface{}{"userID": claims["userID"]}, nil
	}

	return nil, fmt.Errorf("tokenizer parser error")
}
