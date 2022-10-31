package userscontroller

import (
	"log"

	"github.com/Puyodead1/fosscord-server-go/models"
	"github.com/gin-gonic/gin"
)

type GuildAffinity struct {
	GuildID  string `json:"guild_id"`
	Affinity int    `json:"affinity"`
}

type UserAffinity struct {
	UserID   string `json:"user_id"`
	Affinity int    `json:"affinity"`
}

func GetGuildAffinities(c *gin.Context) {
	// id := c.Param("id")

	// user := userservices.GetUserById(id)
	// if user.ID == "" {
	// 	c.JSON(400, fcerrors.HTTPError{Code: jsonerrors.UnknownUser.Code(), Message: jsonerrors.JSONErrorMessages[jsonerrors.UnknownUser]})
	// 	return
	// }
	// get use rfrom content and cast
	user := c.MustGet("CurrentUser").(models.User)
	log.Println(user.ID)

	c.JSON(200, gin.H{"guild_affinities": []GuildAffinity{}})
}

func GetUserAffinities(c *gin.Context) {
	// id := c.Param("id")

	// user := userservices.GetUserById(id)
	// if user.ID == "" {
	// 	c.JSON(400, fcerrors.HTTPError{Code: jsonerrors.UnknownUser.Code(), Message: jsonerrors.JSONErrorMessages[jsonerrors.UnknownUser]})
	// 	return
	// }
	user := c.MustGet("CurrentUser").(models.User)
	log.Println(user.ID)

	c.JSON(200, gin.H{"user_affinities": []UserAffinity{}})
}

func GetLibrary(c *gin.Context) {
	// id := c.Param("id")

	// user := userservices.GetUserById(id)
	// if user.ID == "" {
	// 	c.JSON(400, fcerrors.HTTPError{Code: jsonerrors.UnknownUser.Code(), Message: jsonerrors.JSONErrorMessages[jsonerrors.UnknownUser]})
	// 	return
	// }

	user := c.MustGet("CurrentUser").(models.User)
	log.Println(user.ID)

	empty := make([]interface{}, 0)
	c.JSON(200, empty)
}

func GetBillingLocalizedPricingPromo(c *gin.Context) {
	// id := c.Param("id")

	// user := userservices.GetUserById(id)
	// if user.ID == "" {
	// 	c.JSON(400, fcerrors.HTTPError{Code: jsonerrors.UnknownUser.Code(), Message: jsonerrors.JSONErrorMessages[jsonerrors.UnknownUser]})
	// 	return
	// }
	user := c.MustGet("CurrentUser").(models.User)
	log.Println(user.ID)

	// TODO: country code

	c.JSON(200, gin.H{"country_code": "US", "localized_pricing_promo": nil})
}

func GetBillingPaymentSources(c *gin.Context) {
	// id := c.Param("id")

	// user := userservices.GetUserById(id)
	// if user.ID == "" {
	// 	c.JSON(400, fcerrors.HTTPError{Code: jsonerrors.UnknownUser.Code(), Message: jsonerrors.JSONErrorMessages[jsonerrors.UnknownUser]})
	// 	return
	// }
	user := c.MustGet("CurrentUser").(models.User)
	log.Println(user.ID)

	empty := make([]interface{}, 0)
	c.JSON(200, empty)
}

func GetBillingCountryCode(c *gin.Context) {
	// id := c.Param("id")

	// user := userservices.GetUserById(id)
	// if user.ID == "" {
	// 	c.JSON(400, fcerrors.HTTPError{Code: jsonerrors.UnknownUser.Code(), Message: jsonerrors.JSONErrorMessages[jsonerrors.UnknownUser]})
	// 	return
	// }
	user := c.MustGet("CurrentUser").(models.User)
	log.Println(user.ID)

	// TODO: country code

	c.JSON(200, gin.H{"country_code": "US"})
}
