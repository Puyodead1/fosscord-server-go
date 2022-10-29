package userscontroller

import (
	"github.com/Puyodead1/fosscord-server-go/models"
	userservices "github.com/Puyodead1/fosscord-server-go/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	CaptchaKey            string `json:"captcha_key"`
	Consent               bool   `json:"consent"`
	DOB                   string `json:"dob"`
	Email                 string `json:"email"`
	GiftCodeSkuId         string `json:"gift_code_sku_id"`
	Invite                string `json:"invite"`
	Password              string `json:"password"`
	PromotionalEmailOptIn bool   `json:"promotional_email_opt_in"`
	Username              string `json:"username"`
}

// TODO: captcha
/*
	POST /api/v9/auth/register
	{
		"captcha_key": ["captcha-required"],
		"captcha_sitekey": string,
		"captcha_service": "hcaptcha", // or "recaptcha"?
	}
*/
func Register(c *gin.Context) {
	// validate the request body is RegisterRequest
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	// TODO: proper error responses
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// create a user from the request body
	user := models.User{
		ID:            userservices.GenerateID(),
		Username:      req.Username,
		Email:         &req.Email,
		Password:      string(hashedPassword),
		Discriminator: userservices.GenerateDiscriminator(),
	}

	// create the user in the database
	if err := userservices.CreateUser(&user); err != nil {
		// TODO: proper error responses
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// generate a token
	token, err := userservices.GenerateToken(user.ID)
	if err != nil {
		// TODO: proper error responses
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
