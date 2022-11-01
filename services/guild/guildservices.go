package guildservices

import (
	"github.com/Puyodead1/fosscord-server-go/initializers"
	"github.com/Puyodead1/fosscord-server-go/models"
)

// handles creating a guild in the database
func CreateGuild(guild *models.Guild) error {
	tx := initializers.DB.Create(&guild)

	return tx.Error
}

// handles getting a guild from the database by id
func GetGuild(guild_id string) models.Guild {
	var guild models.Guild
	initializers.DB.Preload("Members").Preload("Channels").Where("id = ?", guild_id).First(&guild)
	return guild
}

// handles updating a guild
func UpdateGuild(guild *models.Guild) error {
	tx := initializers.DB.Save(&guild)

	return tx.Error
}
