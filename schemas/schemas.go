package schemas

type ChannelCreateSchema struct {
}

type GuildCreateSchema struct {
	Name              string         `json:"name" binding:"required,ascii"`
	GuildTemplateCode *string        `json:"guild_template_code"`
	Icon              *string        `json:"icon"`
	SystemChannelID   *string        `json:"system_channel_id"`
	Channels          []*interface{} `json:"channels"` // TODO: struct
}
