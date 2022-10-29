package userscontroller

import (
	"github.com/Puyodead1/fosscord-server-go/models"
	userservices "github.com/Puyodead1/fosscord-server-go/services"
	"github.com/gin-gonic/gin"
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

// TODO: captcha required
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

	// create a user from the request body
	user := models.User{
		ID:            userservices.GenerateID(),
		Username:      req.Username,
		Email:         &req.Email,
		Password:      req.Password, // TODO: hash the password
		Discriminator: userservices.GenerateDiscriminator(),
	}

	// create the user in the database
	if err := userservices.CreateUser(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// TODO: return token
	// return the user
	c.JSON(200, user)
}
