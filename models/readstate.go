package models

type ReadState struct {
	ID                  string  `json:"id" gorm:"primaryKey;unique"`
	ChannelID           string  `json:"channel_id" gorm:"type:string;not null;unique"`
	Channel             Channel `gorm:"many2many:channels;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ChannelID;references:ID"`
	UserID              string  `json:"user_id" gorm:"type:string;not null;unique"`
	User                User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID"`
	LastMessageID       *string `json:"last_message_id" gorm:"type:string"`
	PublicAck           *string `json:"public_ack" gorm:"type:string"`
	NotificationsCursor *int    `json:"notifications_cursor" gorm:"type:int"`
	LastPinTimestamp    *int    `json:"last_pin_timestamp" gorm:"type:int"`
	MentionCount        int     `json:"mention_count" gorm:"type:int;not null;default=0"`
}
