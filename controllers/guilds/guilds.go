package guildscontroller

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/Puyodead1/fosscord-server-go/models"
	"github.com/Puyodead1/fosscord-server-go/schemas"
	guildservices "github.com/Puyodead1/fosscord-server-go/services/guild"
	memberservices "github.com/Puyodead1/fosscord-server-go/services/member"
	"github.com/Puyodead1/fosscord-server-go/utils"
	fcerrors "github.com/Puyodead1/fosscord-server-go/utils/errors"
	"github.com/Puyodead1/fosscord-server-go/utils/errors/fielderror"
	"github.com/Puyodead1/fosscord-server-go/utils/errors/jsonerrors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateGuild(c *gin.Context) {
	var req schemas.GuildCreateSchema
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

	// TODO: rights
	// TODO: check the users max guilds

	guild := models.Guild{
		ID:          utils.GenerateID(),
		Name:        req.Name,
		OwnerID:     c.GetString("CurrentUserID"),
		Channels:    make([]models.Channel, 0),
		Bans:        make([]interface{}, 0),
		Roles:       make([]interface{}, 0),
		Emojis:      make([]interface{}, 0),
		Stickers:    make([]interface{}, 0),
		Invites:     make([]interface{}, 0),
		VoiceStates: make([]interface{}, 0),
		Webhooks:    make([]interface{}, 0),
	}
	if req.Icon != nil {
		guild.Icon = req.Icon
	}

	// create guild
	err := guildservices.CreateGuild(&guild)
	if err != nil {
		c.JSON(500, fcerrors.HTTPError{
			Code:    int(jsonerrors.GeneralError),
			Message: jsonerrors.JSONErrorMessages[jsonerrors.GeneralError],
		})
		return
	}

	// TODO: template
	// TODO: create channels from body

	// TODO: auto join
	err = memberservices.AddToGuild(c.GetString("CurrentUserID"), guild.ID)
	if err != nil {
		log.Println(err)
		c.JSON(500, fcerrors.HTTPError{
			Code:    int(jsonerrors.GeneralError),
			Message: jsonerrors.JSONErrorMessages[jsonerrors.GeneralError],
		})
		return
	}

	c.JSON(201, gin.H{"id": guild.ID})
}
