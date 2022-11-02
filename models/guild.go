package models

type Guild struct {
	ID           string   `json:"id" gorm:"primaryKey;unique"`
	AFKChannelID *string  `json:"afk_channel_id" gorm:"type:string"`
	AFKChannel   *Channel `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AFKChannelID;references:ID"`
	AFKTimeout   int      `json:"afk_timeout" gorm:"type:int;not null;default=3600"`
	// Bans                        []interface{} `json:"bans" gorm:"type:json;default:'[]'"`
	Banner                      *string  `json:"banner" gorm:"type:string"`
	DefaultMessageNotifications int      `json:"default_message_notifications" gorm:"type:int;not null;default=0"`
	Description                 *string  `json:"description" gorm:"type:string"`
	DiscoverySplash             *string  `json:"discovery_splash" gorm:"type:string"`
	ExplicitContentFilter       int      `json:"explicit_content_filter" gorm:"type:int;not null;default=0"`
	Features                    []string `json:"features"  gorm:"type:json;default:'[]'"` // TODO: enum
	PrimaryCategoryID           *string  `json:"primary_category_id" gorm:"type:string"`
	Icon                        *string  `json:"icon" gorm:"type:string"`
	Large                       bool     `json:"large" gorm:"type:boolean;not null;default=false"`
	MaxMembers                  int      `json:"max_members" gorm:"type:int;not null;default=100"`
	MaxPresences                int      `json:"max_presences" gorm:"type:int;not null;default=50000"`
	MaxVideoChannelUsers        int      `json:"max_video_channel_users" gorm:"type:int;not null;default=25"`
	Region                      string   `json:"region" gorm:"type:string;not null;default='us-central'"`
	MemberCount                 int      `json:"member_count" gorm:"type:int;not null;default=0"`
	Members                     []Member `json:"members" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	PresenceCount               int      `json:"presence_count" gorm:"type:int;not null;default=0"`
	// Roles                       []interface{} `json:"roles"  gorm:"type:json;default:'[]'"` // TODO: struct
	Channels   []Channel `json:"channels" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	TemplateID *string   `json:"template_id" gorm:"type:string"`
	// Template                    *interface{}   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID"` // TODO: struct
	// Emojis            []interface{} `json:"emojis" gorm:"type:json;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID;default:'[]'"`       // TODO: struct
	// Stickers          []interface{} `json:"stickers" gorm:"type:json;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID;default:'[]'"`     // TODO: struct
	// Invites           []interface{} `json:"invites" gorm:"type:json;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID;default:'[]'"`      // TODO: struct
	// VoiceStates       []interface{} `json:"voice_states" gorm:"type:json;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID;default:'[]'"` // TODO: struct
	// Webhooks          []interface{} `json:"webooks" gorm:"type:json;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID;default:'[]'"`      // TODO: struct
	MFALevel          int      `json:"mfa_level" gorm:"type:int;not null;default=0"`
	VerificationLevel int      `json:"verification_level" gorm:"type:int;not null;default=0"`
	PreferredLocale   string   `json:"preferred_locale" gorm:"type:string;not null;default='en-US'"`
	Name              string   `json:"name" gorm:"type:string;not null"`
	OwnerID           string   `json:"owner_id" gorm:"type:string;not null"`
	Owner             *User    `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OwnerID;references:ID"`
	SystemChannelID   *string  `json:"system_channel_id" gorm:"type:string"`
	SystemChannel     *Channel `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:SystemChannelID;references:ID"`
	ScheduledEvents   []string `json:"guild_scheduled_events" gorm:"type:json;default:'[]'"`
}
