package userscontroller

import (
	"log"

	"github.com/Puyodead1/fosscord-server-go/models"
	userservices "github.com/Puyodead1/fosscord-server-go/services"
	"github.com/Puyodead1/fosscord-server-go/utils/errors"
	"github.com/Puyodead1/fosscord-server-go/utils/errors/fielderror"
	"github.com/Puyodead1/fosscord-server-go/utils/errors/httperror"
	"github.com/Puyodead1/fosscord-server-go/utils/errors/jsonerrors"
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
// TODO: check if email is a valid email
// TODO: password policy
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

	// check if the email already exists
	if userservices.GetUserByEmail(req.Email).ID != "" {
		c.JSON(400, errors.HTTPError{
			Code:    int(jsonerrors.InvalidFormBody),
			Message: jsonerrors.JSONErrorMessages[jsonerrors.InvalidFormBody],
			Errors: &map[string]errors.FieldError{
				"email": {
					EErrors: []errors.FieldErrorErrors{
						{
							Code:    fielderror.EMAIL_ALREADY_REGISTERED.String(),
							Message: fielderror.EMAIL_ALREADY_REGISTERED.Message(),
						},
					},
				},
			},
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("[Registration] Failed to hash password: %v", err)
		c.JSON(500, errors.HTTPError{Code: 500, Message: string(httperror.InternalServerError)})
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
		log.Fatalf("[Registration] Failed to create user: %v", err)
		c.JSON(500, errors.HTTPError{Code: 500, Message: string(httperror.InternalServerError)})
		return
	}

	// generate a token
	token, err := userservices.GenerateToken(user.ID)
	if err != nil {
		log.Fatalf("[Registration] Failed to generate user: %v", err)
		c.JSON(500, errors.HTTPError{Code: 500, Message: string(httperror.InternalServerError)})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
