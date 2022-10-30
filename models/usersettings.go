package models

type CustomStatus struct {
	Text      string  `json:"text" gorm:"type:string;not null"`
	EmojiID   *string `json:"emoji_id" gorm:"type:string"`
	EmojiName *string `json:"emoji_name" gorm:"type:string"`
	ExpiresAt *int64  `json:"expires_at" gorm:"type:bigint"` // is this the right type?
}

type FriendSourceFlags struct {
	MutualFriends bool `json:"mutual_friends" gorm:"type:bool;not null;default:false"`
	MutualGuilds  bool `json:"mutual_guilds" gorm:"type:bool;not null;default:false"`
}

type GuildFolder struct {
	ID       string `json:"id" gorm:"type:string"`
	Name     string `json:"name" gorm:"type:string"`
	Color    int    `json:"color" gorm:"type:int"` // is this the right type?
	GuildIDs string `json:"guild_ids" gorm:"type:string[];not null;default:'[]'"`
}

type UserSettings struct {
	ID                                string             `json:"id" gorm:"primaryKey"`
	ActivityJoiningRestrictedGuildIDs string             `json:"activity_joining_restricted_guild_ids" gorm:"type:string[];not null;default:'[]'"`
	ActivityRestrictedGuildIDs        string             `json:"activity_restricted_guild_ids" gorm:"type:string[];not null;default:'[]'"`
	AFKTimeout                        int                `json:"afk_timeout" gorm:"type:int;default:3600"`
	AllowAccessibilityDetection       bool               `json:"allow_accessibility_detection" gorm:"type:bool;not null;default:false"`
	AnimateEmoji                      bool               `json:"animate_emoji" gorm:"type:bool;not null;default:true"`
	AnimateStickers                   int                `json:"animate_stickers" gorm:"type:int;not null;default:0"`
	ContactSyncEnabled                bool               `json:"contact_sync_enabled" gorm:"type:bool;not null;default:false"`
	ConvertEmoticons                  bool               `json:"convert_emoticons" gorm:"type:bool;not null;default:true"`
	CustomStatus                      *CustomStatus      `json:"custom_status" gorm:"type:jsonb"` // is this type correct?
	DefaultGuildsRestricted           bool               `json:"default_guilds_restricted" gorm:"type:bool;not null;default:false"`
	DetectPlatformAccounts            bool               `json:"detect_platform_accounts" gorm:"type:bool;not null;default:false"`
	DeveloperMode                     bool               `json:"developer_mode" gorm:"type:bool;not null;default:false"`
	DisableGamesTab                   bool               `json:"disable_games_tab" gorm:"type:bool;not null;default:true"`
	EnableTTSCommand                  bool               `json:"enable_tts_command" gorm:"type:bool;not null;default:true"`
	ExplicitContentFilter             int                `json:"explicit_content_filter" gorm:"type:int;not null;default:0"` // TODO: probably should make an enum for this, but idk what the values are
	FriendDiscoveryFlags              int                `json:"friend_discovery_flags" gorm:"type:int;not null;default:0"`
	FriendSourceFlags                 *FriendSourceFlags `json:"friend_source_flags" gorm:"type:jsonb"`
	GifAutoPlay                       bool               `json:"gif_auto_play" gorm:"type:bool;not null;default:true"`
	GuildFolders                      []GuildFolder      `json:"guild_folders" gorm:"type:jsonb"`
	InlineAttachmentMedia             bool               `json:"inline_attachment_media" gorm:"type:bool;not null;default:true"`
	InlineEmbedMedia                  bool               `json:"inline_embed_media" gorm:"type:bool;not null;default:true"`
	Locale                            string             `json:"locale" gorm:"type:string;not null;default:'en-US'"` // TODO: enum?
	MessageDisplayCompact             bool               `json:"message_display_compact" gorm:"type:bool;not null;default:false"`
	NativePhoneIntegrationEnabled     bool               `json:"native_phone_integration_enabled" gorm:"type:bool;not null;default:false"`
	Passwordless                      bool               `json:"passwordless" gorm:"type:bool;not null;default:false"`
	RenderEmbeds                      bool               `json:"render_embeds" gorm:"type:bool;not null;default:true"`
	RenderReactions                   bool               `json:"render_reactions" gorm:"type:bool;not null;default:true"`
	RestrictedGuilds                  string             `json:"restricted_guilds" gorm:"type:string[];not null;default:'[]'"`
	ShowCurrentGame                   bool               `json:"show_current_game" gorm:"type:bool;not null;default:true"`
	Status                            string             `json:"status" gorm:"type:string;not null;default:'online'"` // TODO: enum
	StreamNotificationsEnabled        bool               `json:"stream_notifications_enabled" gorm:"type:bool;not null;default:false"`
	Theme                             string             `json:"theme" gorm:"type:string;not null;default:'light'"` // TODO: enum?
	TimezoneOffset                    int                `json:"timezone_offset" gorm:"type:int;not null;default:0"`
	ViewNsfwCommands                  bool               `json:"view_nsfw_commands" gorm:"type:bool;not null;default:false"`
	ViewNsfwGuilds                    bool               `json:"view_nsfw_guilds" gorm:"type:bool;not null;default:false"`
}
