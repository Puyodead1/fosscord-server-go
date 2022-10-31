package models

type User struct {
	ID                string        `json:"id" gorm:"primaryKey;unique"`
	Username          string        `json:"username" gorm:"type:string;not null"`
	Avatar            *string       `json:"avatar" gorm:"type:string"`
	AvatarDecoration  *string       `json:"avatar_decoration" gorm:"type:string"` // base64 encoded image
	Discriminator     string        `json:"discriminator" gorm:"type:string;not null"`
	PublicFlags       int           `json:"public_flags" gorm:"type:int;not null;default:0"`
	Flags             int           `json:"flags" gorm:"type:int;not null;default:0"`
	PurchasedFlags    int           `json:"purchased_flags" gorm:"type:int;not null;default:0"`
	PremiumUsageFlags int           `json:"premium_usage_flags" gorm:"type:int;not null;default:0"`
	Banner            *string       `json:"banner" gorm:"type:string"`
	BannerColor       *int          `json:"banner_color" gorm:"type:int"`
	AccentColor       *int          `json:"accent_color" gorm:"type:int"`
	Bio               *string       `json:"bio" gorm:"type:string"`
	NsfwAllowed       bool          `json:"nsfw_allowed" gorm:"type:bool;not null;default:false"`
	MfaEnabled        bool          `json:"mfa_enabled" gorm:"type:bool;not null;default:false"`
	PremiumType       int           `json:"premium_type" gorm:"type:int;not null;default:0"`
	Email             *string       `json:"email" gorm:"type:string"` // email should be nullable since you can create a temp account with an invite
	Verified          bool          `json:"verified" gorm:"type:bool;not null;default:false"`
	Phone             *string       `json:"phone" gorm:"type:string"`
	Password          string        `json:"password" gorm:"type:string;not null"`
	Settings          *UserSettings `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ID;references:ID"`
}
