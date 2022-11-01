package models

type Member struct {
	Index   int     `json:"-" gorm:"primaryKey;autoIncrement"`
	ID      string  `json:"id" gorm:"type:string;not null"` // TODO: unique index for both guild id and user id
	User    User    `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID;references:ID"`
	GuildID string  `json:"guild_id" gorm:"type:string;not null"`
	Guild   Guild   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:GuildID;references:ID"`
	Nick    *string `json:"nick" gorm:"type:string"`
	// Roles []string `json:"roles" gorm:"type:string"` // TODO:
	JoinedAt     string  `json:"joined_at" gorm:"type:string;not null"`
	PremiumSince *string `json:"premium_since" gorm:"type:string"`
	Deaf         bool    `json:"deaf" gorm:"type:boolean;not null;default=false"`
	Mute         bool    `json:"mute" gorm:"type:boolean;not null;default=false"`
	Pending      bool    `json:"pending" gorm:"type:boolean;not null;default=false"` // TODO: what is this?
	// Settings UserGuildSettings `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID;references:ID"` // TODO:
	LastMessageID               *string `json:"last_message_id" gorm:"type:string"`
	Avatar                      *string `json:"avatar" gorm:"type:string"`
	Banner                      *string `json:"banner" gorm:"type:string"`
	Bio                         *string `json:"bio" gorm:"type:string"`
	CommunicationDisabledUntil  *string `json:"communication_disabled_until" gorm:"type:string"`
	Flags                       int     `json:"flags" gorm:"type:integer;not null;default=0"`
	PopoutAnimationParticleType string  `json:"popout_animation_particle_type" gorm:"type:string"` // TODO: what is this?
	ThemeColors                 []int   `json:"theme_colors" gorm:"type:string"`                   // array of ints
}
