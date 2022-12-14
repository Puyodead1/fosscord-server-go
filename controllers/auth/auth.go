package authcontroller

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/Puyodead1/fosscord-server-go/models"
	jwtservices "github.com/Puyodead1/fosscord-server-go/services/jwt"
	userservices "github.com/Puyodead1/fosscord-server-go/services/user"
	"github.com/Puyodead1/fosscord-server-go/utils"
	fcerrors "github.com/Puyodead1/fosscord-server-go/utils/errors"
	"github.com/Puyodead1/fosscord-server-go/utils/errors/fielderror"
	"github.com/Puyodead1/fosscord-server-go/utils/errors/httperror"
	"github.com/Puyodead1/fosscord-server-go/utils/errors/jsonerrors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// TODO: middleware for json parsing?

// TODO: how do we validate boolean values? using binding="boolean" results in validation errors as if they weren't even provided
type RegisterRequest struct {
	CaptchaKey            string `json:"captcha_key,omitempty" binding:"ascii"`
	Consent               bool   `json:"consent"`
	DOB                   string `json:"date_of_birth" binding:"required,ascii"`
	Email                 string `json:"email" binding:"required,email"`
	GiftCodeSkuId         string `json:"gift_code_sku_id,omitempty" binding:"ascii"`
	Invite                string `json:"invite,omitempty" binding:"ascii"`
	Password              string `json:"password" binding:"required,printascii,min=6,max=72"` // TODO: password policy
	PromotionalEmailOptIn bool   `json:"promotional_email_opt_in"`
	Username              string `json:"username" binding:"required,ascii,min=2,max=32"`
}

type LoginRequest struct {
	Login         string `json:"login" binding:"required,ascii"` // Email or phonenumber
	Password      string `json:"password" binding:"required"`
	Undelete      bool   `json:"undelete"`
	CaptchaKey    string `json:"captcha_key,omitempty" binding:"ascii"`
	LoginSource   any    `json:"login_source,omitempty"` // TODO: what is the type for this?
	GiftCodeSkuId string `json:"gift_code_sku_id,omitempty" binding:"ascii"`
}

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
		var serr *json.SyntaxError
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			e := fcerrors.HTTPError{
				Code:    int(jsonerrors.InvalidFormBody),
				Message: jsonerrors.JSONErrorMessages[jsonerrors.InvalidFormBody],
			}
			errs := err.(validator.ValidationErrors)
			for _, err := range errs {
				e.Errors = &map[string]fcerrors.FieldError{
					err.Field(): {
						EErrors: []fcerrors.FieldErrorErrors{
							{
								Code:    fielderror.ValidationErrors[err.Tag()].String(),
								Message: fielderror.ValidationErrors[err.Tag()].Message(),
							},
						},
					},
				}
			}
			c.JSON(400, e)
			return
		} else if errors.As(err, &serr) {
			c.JSON(400, fcerrors.HTTPError{
				Code:    int(jsonerrors.InvalidJSON),
				Message: jsonerrors.JSONErrorMessages[jsonerrors.InvalidJSON],
			})
			return
		} else {
			c.JSON(400, fcerrors.HTTPError{
				Code:    int(jsonerrors.GeneralError),
				Message: jsonerrors.JSONErrorMessages[jsonerrors.GeneralError],
			})
			return
		}
	}

	// TODO: consent check
	// TODO: captcha
	// TODO: registration disabled check
	// TODO: block proxies
	// TODO: password policy
	// TODO: guest/temp accounts when no password provided

	// check if the email already exists
	if userservices.GetUserByEmail(req.Email).ID != "" {
		c.JSON(400, fcerrors.HTTPError{
			Code:    int(jsonerrors.InvalidFormBody),
			Message: jsonerrors.JSONErrorMessages[jsonerrors.InvalidFormBody],
			Errors: &map[string]fcerrors.FieldError{
				"email": {
					EErrors: []fcerrors.FieldErrorErrors{
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

	// TODO: underage check
	// TODO: ratelimit
	// TODO: add to guild if registered with invite

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("[Registration] Failed to hash password: %v", err)
		c.JSON(500, fcerrors.HTTPError{Code: 500, Message: string(httperror.InternalServerError)})
		return
	}

	// create a user from the request body
	usersettings := models.UserSettings{}
	user := models.User{
		ID:            utils.GenerateID(),
		Username:      req.Username,
		Email:         &req.Email,
		Password:      string(hashedPassword),
		Discriminator: userservices.GenerateDiscriminator(),
		Settings:      &usersettings,
	}

	// create the user in the database
	if err := userservices.CreateUser(&user); err != nil {
		log.Fatalf("[Registration] Failed to create user: %v", err)
		c.JSON(500, fcerrors.HTTPError{Code: 500, Message: string(httperror.InternalServerError)})
		return
	}

	// generate a token
	token, err := jwtservices.GenerateToken(user.ID)
	if err != nil {
		log.Fatalf("[Registration] Failed to generate user: %v", err)
		c.JSON(500, fcerrors.HTTPError{Code: 500, Message: string(httperror.InternalServerError)})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func Login(c *gin.Context) {
	// validate the request body is LoginRequest
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		var serr *json.SyntaxError
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			e := fcerrors.HTTPError{
				Code:    int(jsonerrors.InvalidFormBody),
				Message: jsonerrors.JSONErrorMessages[jsonerrors.InvalidFormBody],
			}
			errs := err.(validator.ValidationErrors)
			for _, err := range errs {
				e.Errors = &map[string]fcerrors.FieldError{
					err.Field(): {
						EErrors: []fcerrors.FieldErrorErrors{
							{
								Code:    fielderror.ValidationErrors[err.Tag()].String(),
								Message: fielderror.ValidationErrors[err.Tag()].Message(),
							},
						},
					},
				}
			}
			c.JSON(400, e)
			return
		} else if errors.As(err, &serr) {
			c.JSON(400, fcerrors.HTTPError{
				Code:    int(jsonerrors.InvalidJSON),
				Message: jsonerrors.JSONErrorMessages[jsonerrors.InvalidJSON],
			})
			return
		} else {
			c.JSON(400, fcerrors.HTTPError{
				Code:    int(jsonerrors.GeneralError),
				Message: jsonerrors.JSONErrorMessages[jsonerrors.GeneralError],
			})
			return
		}
	}

	// TODO: captcha

	// check if user exists
	user := userservices.GetUserByLogin(req.Login)
	if user.ID == "" {
		c.JSON(400, fcerrors.HTTPError{
			Code:    int(jsonerrors.InvalidFormBody),
			Message: jsonerrors.JSONErrorMessages[jsonerrors.InvalidFormBody],
			Errors: &map[string]fcerrors.FieldError{
				"login": {
					EErrors: []fcerrors.FieldErrorErrors{
						{
							Code:    fielderror.INVALID_LOGIN.String(),
							Message: fielderror.INVALID_LOGIN.Message(),
						},
					},
				},
				"password": {
					EErrors: []fcerrors.FieldErrorErrors{
						{
							Code:    fielderror.INVALID_LOGIN.String(),
							Message: fielderror.INVALID_LOGIN.Message(),
						},
					},
				},
			},
		})
		return
	}

	// TODO: undelete

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(400, fcerrors.HTTPError{
			Code:    int(jsonerrors.InvalidFormBody),
			Message: jsonerrors.JSONErrorMessages[jsonerrors.InvalidFormBody],
			Errors: &map[string]fcerrors.FieldError{
				"login": {
					EErrors: []fcerrors.FieldErrorErrors{
						{
							Code:    fielderror.INVALID_LOGIN.String(),
							Message: fielderror.INVALID_LOGIN.Message(),
						},
					},
				},
				"password": {
					EErrors: []fcerrors.FieldErrorErrors{
						{
							Code:    fielderror.INVALID_LOGIN.String(),
							Message: fielderror.INVALID_LOGIN.Message(),
						},
					},
				},
			},
		})
		return
	}

	// TODO: new location detected
	// TODO: mfa

	// generate a token
	token, err := jwtservices.GenerateToken(user.ID)
	if err != nil {
		log.Fatalf("[Registration] Failed to generate user: %v", err)
		c.JSON(500, fcerrors.HTTPError{Code: 500, Message: string(httperror.InternalServerError)})
		return
	}

	// TODO: user settings
	c.JSON(200, gin.H{"token": token, "user_id": user.ID, "user_settings": user.Settings})
}
