package userservices

import (
	"fmt"
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

// gets a user by their id
func GetUserById(id string) models.User {
	var user models.User
	initializers.DB.Preload("Settings").First(&user, id)
	return user
}

// gets a a user by their email
func GetUserByEmail(email string) models.User {
	var user models.User
	initializers.DB.Preload("Settings").Where("email = ?", email).First(&user)
	return user
}

// gets a user by their phone number
func GetUserByPhone(phone string) models.User {
	var user models.User
	initializers.DB.Preload("Settings").Where("phone = ?", phone).First(&user)
	return user
}

// gets a user by their email or phone number
func GetUserByLogin(q string) models.User {
	var user models.User
	initializers.DB.Preload("Settings").Where("email = ? OR phone = ?", q, q).First(&user) // TODO: omit the id column
	return user
}

// gets all users
func GetUsers() []models.User {
	var users []models.User
	initializers.DB.Preload("Settings").Find(&users)
	return users
}

// updates a user
func UpdateUser(id string, user *models.User) {
	initializers.DB.Model(&user).Preload("Settings").Where("id = ?", id).Updates(user)
}

// delets a user
func DeleteUser(id string) {
	var user models.User
	initializers.DB.Preload("Settings").Delete(&user, id)
}

// get user settings
func GetUserSettings(id string) models.UserSettings {
	var settings models.UserSettings
	initializers.DB.First(&settings, id)
	return settings
}

// get read states for user
func GetReadStates(id string) []models.ReadState {
	var readStates []models.ReadState
	initializers.DB.Where("user_id = ?", id).Find(&readStates)
	return readStates
}

// generates a random discriminator
func GenerateDiscriminator() string {
	return strconv.Itoa(0001 + rand.Intn(9999-0001))
}

// generates a snowflake id
func GenerateID() string {
	return initializers.Node.Generate().String()
}

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

func GenerateSessionID() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", key)
}
