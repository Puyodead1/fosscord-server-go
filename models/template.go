package models

type Template struct {
	Code                  string  `json:"code" gorm:"primaryKey;unique"`
	Name                  string  `json:"name" gorm:"type:string;not null"`
	Description           *string `json:"description" gorm:"type:string"`
	UsageCount            int     `json:"usage_count" gorm:"type:integer;not null;default=0"`
	CreatorID             string  `json:"creator_id" gorm:"type:string;not null"`
	Creator               User    `json:"creator" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CreatorID;references:ID"`
	CreatedAt             string  `json:"created_at" gorm:"type:string;not null"`
	UpdatedAt             string  `json:"updated_at" gorm:"type:string;not null"`
	SourceGuildID         string  `json:"source_guild_id" gorm:"type:string;not null"`
	SerializedSourceGuild Guild   `json:"serialized_source_guild" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:SourceGuildID;references:ID"`
	IsDirty               bool    `json:"is_dirty" gorm:"type:boolean;not null;default=false"`
}
