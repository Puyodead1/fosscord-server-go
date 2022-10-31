package models

type Guild struct {
	ID           string   `json:"id" gorm:"primaryKey;unique"`
	AFKChannelID *string  `json:"afk_channel_id" gorm:"type:string"`
	AFKChannel   *Channel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AFKChannelID;references:ID"`
	AFKTimeout   int      `json:"afk_timeout" gorm:"type:int;not null;default=3600"`
	// Bans                        []*interface{} // TODO: struct
	Banner                      *string  `json:"banner" gorm:"type:string"`
	DefaultMessageNotifications int      `json:"default_message_notifications" gorm:"type:int;not null;default=0"`
	Description                 *string  `json:"description" gorm:"type:string"`
	DiscoverySplash             *string  `json:"discovery_splash" gorm:"type:string"`
	ExplicitContentFilter       int      `json:"explicit_content_filter" gorm:"type:int;not null;default=0"`
	Features                    []string `json:"features" gorm:"type:string"` // TODO: enum
	PrimaryCategoryID           *string  `json:"primary_category_id" gorm:"type:string"`
	Icon                        *string  `json:"icon" gorm:"type:string"`
	Large                       bool     `json:"large" gorm:"type:boolean;not null;default=false"`
	MaxMembers                  int      `json:"max_members" gorm:"type:int;not null;default=100"`
	MaxPresences                int      `json:"max_presences" gorm:"type:int;not null;default=50000"`
	MaxVideoChannelUsers        int      `json:"max_video_channel_users" gorm:"type:int;not null;default=25"`
	MemberCount                 int      `json:"member_count" gorm:"type:int;not null;default=0"`
	// Members                     []interface{}  // TODO: struct
	PresenceCount int `json:"presence_count" gorm:"type:int;not null;default=0"`
	// Roles                       []interface{}  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID"` // TODO: struct
	Channels   []Channel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	TemplateID *string   `json:"template_id" gorm:"type:string"`
	// Template                    *interface{}   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID"` // TODO: struct
	// Emojis                      []interface{}  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID"` // TODO: struct
	// Stickers                    []interface{}  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID"` // TODO: struct
	// Invites                     []interface{}  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID"` // TODO: struct
	// VoiceStates                 []interface{}  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID"` // TODO: struct
	// Webhooks                    []interface{}  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID"` // TODO: struct
	MFALevel int    `json:"mfa_level" gorm:"type:int;not null;default=0"`
	Name     string `json:"name" gorm:"type:string;not null"`
	OwnerID  string `json:"owner_id" gorm:"type:string;not null"`
	Owner    *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OwnerID;references:ID"`
}
