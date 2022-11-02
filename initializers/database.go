package initializers

import (
	"errors"
	"log"

	"github.com/Puyodead1/fosscord-server-go/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.UserSettings{})
	DB.AutoMigrate(&models.Channel{})
	DB.AutoMigrate(&models.Guild{})
	DB.AutoMigrate(&models.Member{})
	DB.AutoMigrate(&models.ReadState{})
	DB.AutoMigrate(&models.Template{})

	if err := DB.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		tx := DB.Create(&models.User{
			ID:       "268473310986240001",
			Username: "Fosscord",
			// Avatar:           nil,
			// AvatarDecoration: nil,
			Discriminator: "0001",
			PublicFlags:   131072,
		})

		if tx.Error != nil {
			log.Fatalf("Failed to create default user: %v", tx.Error)
		}
	}

	if err := DB.First(&models.Guild{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		textParent := "000000000000000001"
		voiceParent := "000000000000000003"

		channels := []models.Channel{
			{
				ID:   textParent,
				Type: 4,
				Name: "Text Channels",
				// Topic:                      nil,
				Bitrate: 64000,
				NSFW:    false,
				// ParentID:                   nil,
				// DefaultAutoArchiveDuration: nil,
				PermissionOverwrites: make([]interface{}, 0),
				AvailableTags:        make([]interface{}, 0),
				// DefaultSortOrder:           nil,
				ReadStates: make([]models.ReadState, 0),
				Recipients: make([]interface{}, 0),
			},
			{
				ID:   "000000000000000002",
				Type: 0,
				Name: "general",
				// Topic:                      nil,
				Bitrate:  64000,
				NSFW:     false,
				ParentID: &textParent,
				// DefaultAutoArchiveDuration: nil,
				PermissionOverwrites: make([]interface{}, 0),
				AvailableTags:        make([]interface{}, 0),
				// DefaultSortOrder:           nil,
				ReadStates: make([]models.ReadState, 0),
				Recipients: make([]interface{}, 0),
			},
			{
				ID:   voiceParent,
				Type: 4,
				Name: "Voice Channels",
				// Topic:                      nil,
				Bitrate: 64000,
				NSFW:    false,
				// ParentID:                   nil,
				// DefaultAutoArchiveDuration: nil,
				PermissionOverwrites: make([]interface{}, 0),
				AvailableTags:        make([]interface{}, 0),
				// DefaultSortOrder:           nil,
				ReadStates: make([]models.ReadState, 0),
				Recipients: make([]interface{}, 0),
			},
			{
				ID:   "000000000000000004",
				Type: 2,
				Name: "General",
				// Topic:                      nil,
				Bitrate:  64000,
				NSFW:     false,
				ParentID: &voiceParent,
				// DefaultAutoArchiveDuration: nil,
				PermissionOverwrites: make([]interface{}, 0),
				AvailableTags:        make([]interface{}, 0),
				// DefaultSortOrder:           nil,
				ReadStates: make([]models.ReadState, 0),
				Recipients: make([]interface{}, 0),
			},
		}
		defaultGuild := &models.Guild{
			ID:   "700811170902179862",
			Name: "Blank Server",
			// Description:                 nil,
			Region:                      "us-west",
			VerificationLevel:           0,
			DefaultMessageNotifications: 0,
			ExplicitContentFilter:       0,
			PreferredLocale:             "en-US",
			AFKTimeout:                  300,
			// AFKChannelID:                nil,
			// SystemChannelID:             nil,
			Channels: channels,
			Features: []string{"COMMUNITY"},
			OwnerID:  "268473310986240001",
		}
		tx := DB.Create(defaultGuild)

		if tx.Error != nil {
			log.Fatalf("Failed to create default guild: %v", tx.Error)
		}
	}

	if err := DB.First(&models.Template{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		defaultTemplate := &models.Template{
			Code: "2TffvPucqHkN",
			Name: "Blank Server",
			// Description:   nil,
			UsageCount:    0,
			CreatorID:     "268473310986240001",
			CreatedAt:     "2021-01-01T00:00:00.000000+00:00",
			UpdatedAt:     "2021-01-01T00:00:00.000000+00:00",
			SourceGuildID: "268473310986240001",
			IsDirty:       false,
		}

		tx := DB.Create(defaultTemplate)

		if tx.Error != nil {
			log.Fatalf("Failed to create default template: %v", tx.Error)
		}
	}
}
