package jwtservices

import (
	"fmt"
	"log"
	"time"

	"github.com/Puyodead1/fosscord-server-go/initializers"
	userservices "github.com/Puyodead1/fosscord-server-go/services/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// generates a jwt token
func GenerateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["iat"] = time.Now().Unix()

	tokenString, err := token.SignedString(initializers.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// verifies a jwt token
func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return initializers.SecretKey, nil
	})
	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	return id, nil
}

func TokenValid(c *gin.Context) error {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(initializers.SecretKey), nil
	})
	if err != nil {
		log.Println(err)
		return err
	}
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	if id == "" {
		return fmt.Errorf("failed to extract id")
	}

	user := userservices.GetUserById(id)
	if user.ID == "" {
		return fmt.Errorf("failed to get user")
	}

	c.Set("CurrentUser", user)
	return nil
}

func ExtractTokenID(c *gin.Context) (string, error) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(initializers.SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid := claims["id"].(string)
		return uid, nil
	}
	return "", nil
}
