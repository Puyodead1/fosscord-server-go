package gateway

import "github.com/Puyodead1/fosscord-server-go/models"

type IdentifyPayloadProperties struct {
	OS                     *string `json:"os"`
	Browser                *string `json:"browser"`
	Device                 *string `json:"device"`
	SystemLocale           *string `json:"system_locale"`
	BrowserUserAgent       *string `json:"browser_user_agent"`
	BrowserVersion         *string `json:"browser_version"`
	OSVersion              *string `json:"os_version"`
	Referrer               *string `json:"referrer"`
	ReferringDomain        *string `json:"referring_domain"`
	ReferrerCurrent        *string `json:"referrer_current"`
	ReferringDomainCurrent *string `json:"referring_domain_current"`
	ReleaseChannel         *string `json:"release_channel"`
	ClientBuildNumber      *int    `json:"client_build_number"`
	ClientEventSource      *int    `json:"client_event_source"`
}

type IdentifyPayloadPresense struct {
	Since      *int                   `json:"since"`
	Status     *string                `json:"status"`     // either "idle", "dnd", "online", or "offline". enum?
	Activities map[string]interface{} `json:"activities"` // TODO: struct
	AFK        *bool                  `json:"afk"`        // whether or not the user is afk
}

type IdentifyPayloadClientState struct {
	GuildHashes              map[string]string `json:"guild_hashes"`
	HighestLastMessageID     string            `json:"highest_last_message_id"`
	ReadStateVersion         int               `json:"read_state_version"`
	UserGuildSettingsVersion int               `json:"user_guild_settings_version"`
	UserSettingsVersion      int               `json:"user_settings_version"`
	PrivateChannelsVersion   int               `json:"private_channels_version"`
}

type IdentifyPayload struct {
	Token       string                     `json:"token" validate:"required,base64"`
	Properties  IdentifyPayloadProperties  `json:"properties"`
	Presence    IdentifyPayloadPresense    `json:"presence"`
	Compress    bool                       `json:"compress" validate:"boolean"`
	ClientState IdentifyPayloadClientState `json:"client_state"`
}

// fuck do we call this?
type ReadyEventDataStruct1 struct {
	Entries []any `json:"entries"`
	Partial bool  `json:"partial"`
	Version int   `json:"version"`
}

type ReadyEventPayload struct {
	V                     int                    `json:"v"`
	User                  models.User            `json:"user"`
	PrivateChannels       []any                  `json:"private_channels"` // TODO:
	SessionID             string                 `json:"session_id"`
	Guilds                []any                  `json:"guilds"` // TODO:
	AnalyticsToken        *string                `json:"analytics_token"`
	ConnectedAccounts     *[]any                 `json:"connected_accounts"` // TODO:
	Consents              *[]any                 `json:"consents"`           // TODO:
	CountryCode           *string                `json:"country_code"`
	FriendSuggestionCount *int                   `json:"friend_suggestion_count"`
	GeoOrderedRtcRegions  *[]string              `json:"geo_ordered_rtc_regions"` // TODO:
	Experiments           *[]any                 `json:"experiments"`             // TODO:
	GuildExperiments      *[]any                 `json:"guild_experiments"`       // TODO:
	GuildJoinRequests     *[]any                 `json:"guild_join_requests"`     // TODO:
	Shard                 *[]int                 `json:"shard"`                   // TODO:
	UserSettings          *models.UserSettings   `json:"user_settings"`
	Relationships         *[]any                 `json:"relationships"`       // TODO:
	ReadState             *ReadyEventDataStruct1 `json:"read_state"`          // TODO:
	UserGuildSettings     *ReadyEventDataStruct1 `json:"user_guild_settings"` // TODO:
	Application           *interface{}           `json:"application"`         // TODO:
	MergedMembers         *[]any                 `json:"merged_members"`      // TODO:
	Users                 *[]any                 `json:"users"`               // TODO:
}
