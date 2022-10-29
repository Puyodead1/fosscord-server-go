package userservices

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/Puyodead1/fosscord-server-go/initializers"
	"github.com/Puyodead1/fosscord-server-go/models"
	"github.com/golang-jwt/jwt/v4"
)

// handles creating a user in the database
func CreateUser(user *models.User) error {
	tx := initializers.DB.Create(&user)

	return tx.Error
}

func GetUser(id string) models.User {
	var user models.User
	initializers.DB.First(&user, id)
	return user
}

func GetUsers() []models.User {
	var users []models.User
	initializers.DB.Find(&users)
	return users
}

func UpdateUser(id string, user *models.User) {
	initializers.DB.Model(&user).Where("id = ?", id).Updates(user)
}

func DeleteUser(id string) {
	var user models.User
	initializers.DB.Delete(&user, id)
}

func GenerateDiscriminator() string {
	return strconv.Itoa(0001 + rand.Intn(9999-0001))
}

func GenerateID() string {
	return initializers.Node.Generate().String()
}

func GenerateToken(id string) (string, error) {
	iat := time.Now().UnixMilli()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["iat"] = iat

	tokenString, err := token.SignedString(initializers.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
