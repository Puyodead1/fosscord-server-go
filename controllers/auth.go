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
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	CaptchaKey            string `json:"captcha_key" binding:"ascii"`
	Consent               bool   `json:"consent" binding:"boolean"`
	DOB                   string `json:"dob" binding:"required,ascii"`
	Email                 string `json:"email" binding:"required,email"`
	GiftCodeSkuId         string `json:"gift_code_sku_id" binding:"ascii"`
	Invite                string `json:"invite" binding:"ascii"`
	Password              string `json:"password" binding:"required,printascii,min=8"` // TODO: password policy
	PromotionalEmailOptIn bool   `json:"promotional_email_opt_in" binding:"boolean"`
	Username              string `json:"username" binding:"required,ascii,min=2,max=32"`
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
		e := errors.HTTPError{
			Code:    int(jsonerrors.InvalidFormBody),
			Message: jsonerrors.JSONErrorMessages[jsonerrors.InvalidFormBody],
		}
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			e.Errors = &map[string]errors.FieldError{
				err.Field(): {
					EErrors: []errors.FieldErrorErrors{
						{
							Code:    err.Tag(),
							Message: err.Error(),
						},
					},
				},
			}
		}
		c.JSON(400, e)
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
