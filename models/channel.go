package models

type ChannelType int

const (
	ChannelType_GUILD_TEXT           ChannelType = iota      // a text channel within a guild
	ChannelType_DM                                           // a direct message between users
	ChannelType_GUILD_VOICE                                  // a voice channel within a guild
	ChannelType_GROUP_DM                                     // a direct message between multiple users
	ChannelType_GUILD_CATEGORY                               // an organizational category that contains zero or more channels
	ChannelType_GUILD_NEWS                                   // a channel that users can follow and crosspost into a guild or route
	ChannelType_GUILD_STORE                                  // a channel in which game developers can sell their things
	ChannelType_ENCRYPTED                                    // end-to-end encrypted channel
	ChannelType_ENCRYPTED_THREAD                             // end-to-end encrypted thread channel
	ChannelType_TRANSACTIONAL                                // event chain style transactional channel
	ChannelType_GUILD_NEWS_THREAD                            // a temporary sub-channel within a GUILD_NEWS channel
	ChannelType_GUILD_PUBLIC_THREAD                          // a temporary sub-channel within a GUILD_TEXT channel
	ChannelType_GUILD_PRIVATE_THREAD                         // a temporary sub-channel within a GUILD_TEXT channel that is only viewable by those invited and those with the MANAGE_THREADS permission
	ChannelType_GUILD_STAGE_VOICE                            // a voice channel for hosting events with an audience
	ChannelType_DIRECTORY                                    // guild directory listing channel
	ChannelType_GUILD_FORUM                                  // forum composed of IM threads
	ChannelType_TICKET_TRACKER       ChannelType = iota + 17 // ticket tracker, individual ticket items shall have type 12
	ChannelType_KANBAN                                       // confluence like kanban board
	ChannelType_VOICELESS_WHITEBOARD                         // whiteboard but without voice (whiteboard + voice is the same as stage)
	ChannelType_CUSTOM_START         ChannelType = iota + 45 // start custom channel types from here
	ChannelType_UNHANDLED            ChannelType = 255       // unhandled unowned pass-through channel type
)

type Channel struct {
	ID                         string        `json:"id" gorm:"primaryKey"`
	CreatedAt                  int           `json:"created_at" gorm:"type:int;not null"`
	Name                       string        `json:"name" gorm:"type:string;not null"`
	Icon                       string        `json:"icon" gorm:"type:string"`
	Type                       ChannelType   `json:"type" gorm:"type:int;not null"`
	Recipients                 []interface{} `json:"recipients"  gorm:"type:json;default:'[]'"`
	LastMessageID              string        `json:"last_message_id" gorm:"type:string;not null"`
	GuildID                    *string       `json:"guild_id" gorm:"type:string"`
	Guild                      Guild         `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:GuildID;references:ID"`
	ParentID                   *string       `json:"parent_id" gorm:"type:string"`
	Parent                     *Channel      `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ParentID;references:ID"`
	OwnerID                    *string       `json:"owner_id" gorm:"type:string"`
	Owner                      User          `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OwnerID;references:ID"`
	LastPinTimestamp           *int          `json:"last_pin_timestamp" gorm:"type:int"`
	DefaultAutoArchiveDuration *int          `json:"default_auto_archive_duration" gorm:"type:int"`
	Position                   int           `json:"position" gorm:"type:int;default:0"`
	PermissionOverwrites       []interface{} `json:"permission_overwrites" gorm:"type:json;default:'[]'"` // TODO: struct
	VideoQualityMode           *int          `json:"video_quality_mode" gorm:"type:int"`
	Bitrate                    int           `json:"bitrate" gorm:"type:int;default:64000"`
	UserLimit                  int           `json:"user_limit" gorm:"type:int;default:0"`
	NSFW                       bool          `json:"nsfw" gorm:"type:bool;default:false"`
	RateLimitPerUser           int           `json:"rate_limit_per_user" gorm:"type:int;default=0"`
	Topic                      *string       `json:"topic" gorm:"type:string"`
	// Invites                       *[]interface{} `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID"` // TODO: struct
	RetentionPolicyID *string `json:"retention_policy_id" gorm:"type:string"`
	// Messages                      []interface{}  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID"`
	// VoiceStates                   []interface{}  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID"` // TODO: struct
	ReadStates []ReadState `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ChannelID"`
	// Webhooks                      []interface{}  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID"` // TODO: struct
	Flags                         int  `json:"flags" gorm:"type:int;not null;default:0"`
	DefaultThreadRateLimitPerUser *int `json:"default_thread_rate_limit_per_user" gorm:"type:int;not null;default:0"`
	// DefaultReactionEmoji 		*[]interface{} `json:"default_reaction_emoji" gorm:"type:jsonb"` // TODO: what is the type
	DefaultSortOrder *int          `json:"default_sort_order" gorm:"type:int"`
	AvailableTags    []interface{} `json:"available_tags"  gorm:"type:json;default:'[]'"`
}
