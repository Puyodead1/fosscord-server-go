package memberservices

import (
	"fmt"
	"log"
	"time"

	"github.com/Puyodead1/fosscord-server-go/gateway"
	"github.com/Puyodead1/fosscord-server-go/initializers"
	"github.com/Puyodead1/fosscord-server-go/models"
	guildservices "github.com/Puyodead1/fosscord-server-go/services/guild"
	userservices "github.com/Puyodead1/fosscord-server-go/services/user"
)

// handles creating a member in the database
func CreateMember(member *models.Member) error {
	tx := initializers.DB.Create(&member)

	return tx.Error
}

func IsInGuild(user_id string, guild_id string) bool {
	guild := guildservices.GetGuild(guild_id)
	if guild.ID == "" {
		return false
	}

	for _, member := range guild.Members {
		if member.ID == user_id {
			return true
		}
	}

	return false
}

func AddToGuild(user_id string, guild_id string) error {
	user := userservices.GetUserById(user_id)
	if user.ID == "" {
		return fmt.Errorf("user not found")
	}
	// TODO: check if user is banned from guild
	// TODO: check max guilds

	guild := guildservices.GetGuild(guild_id)
	if guild.ID == "" {
		return fmt.Errorf("guild not found")
	}
	// guild.Bans = make([]interface{}, 0)
	// guild.Emojis = make([]interface{}, 0)
	// guild.Stickers = make([]interface{}, 0)
	// guild.Invites = make([]interface{}, 0)
	// guild.VoiceStates = make([]interface{}, 0)
	// guild.Webhooks = make([]interface{}, 0)

	if IsInGuild(user_id, guild_id) {
		return fmt.Errorf("user is already in guild")
	}

	member := models.Member{
		ID:      user_id,
		GuildID: guild_id,
		// TODO: @everyone role
		JoinedAt: time.Now().String(),
		Pending:  false,
	}

	guild.MemberCount++
	err := CreateMember(&member)
	if err != nil {
		return err
	}

	gmaPayload := &gateway.GuildMemberAdd{
		GuildID: guild_id,
		Member:  member,
	}
	if err != nil {
		log.Println(err)
		return err
	}
	gateway.EmitEvent(gateway.Event{
		EventName: "GUILD_MEMBER_ADD",
		Data:      gmaPayload,
		GuildID:   &guild_id,
	})

	guild.Members = append(guild.Members, member)
	guild.MemberCount = guild.MemberCount + 1

	err = guildservices.UpdateGuild(&guild)
	if err != nil {
		return err
	}

	gcPayload := &gateway.GuildCreate{
		Guild:    guild,
		JoinedAt: member.JoinedAt,
	}
	gateway.EmitEvent(gateway.Event{
		EventName: "GUILD_CREATE",
		Data:      gcPayload,
		UserID:    &user_id,
	})
	return nil

}
