package jsonerrors

type JSONError int

const (
	GeneralError                             JSONError = 0
	UnknownAccount                           JSONError = 10001
	UnknownApplication                       JSONError = 10002
	UnknownChannel                           JSONError = 10003
	UnknownGuild                             JSONError = 10004
	UnknownIntegration                       JSONError = 10005
	UnknownInvite                            JSONError = 10006
	UnknownMember                            JSONError = 10007
	UnknownMessage                           JSONError = 10008
	UnknownPermissionOverwrite               JSONError = 10009
	UnknownProvider                          JSONError = 10010
	UnknownRole                              JSONError = 10011
	UnknownToken                             JSONError = 10012
	UnknownUser                              JSONError = 10013
	UnknownEmoji                             JSONError = 10014
	UnknownWebhook                           JSONError = 10015
	UnknownWebhookService                    JSONError = 10016
	UnknownSession                           JSONError = 10020
	UnknownBan                               JSONError = 10026
	UnknownSKU                               JSONError = 10027
	UnknownStoreListing                      JSONError = 10028
	UnknownEntitlement                       JSONError = 10029
	UnknownBuild                             JSONError = 10030
	UnknownLobby                             JSONError = 10031
	UnknownBranch                            JSONError = 10032
	UnknownStoreDirectoryLayout              JSONError = 10036
	UnknownRedistributable                   JSONError = 10037
	UnknownGiftCode                          JSONError = 10038
	UnknownStream                            JSONError = 10049
	UnknownPremiumServerSubscriptionCooldown JSONError = 10050
	UnknownGuildTemplate                     JSONError = 10057
	UnknownDiscoveryCategory                 JSONError = 10059
	UnknownSticker                           JSONError = 10060
	UnknownInteraction                       JSONError = 10062
	UnknownApplicationCommand                JSONError = 10063
	UnknownVoiceState                        JSONError = 10065
	UnknownApplicationCommandPermissions     JSONError = 10066
	UnknownStageInstance                     JSONError = 10067
	UnknownGuildMemberVerificationForm       JSONError = 10068
	UnknownGuildWelcomeScreen                JSONError = 10069
	UnknownGuildScheduledEvent               JSONError = 10070
	UnknownGuildScheduledEventUser           JSONError = 10071
	UnknownTag                               JSONError = 10087
	NoBots                                   JSONError = 20001
	OnlyBots                                 JSONError = 20002
	ExplicitContentBlocked                   JSONError = 20009
	ApplicationUnauthorized                  JSONError = 20012
	SlowModeRateLimit                        JSONError = 20016
	AccountOwnerOnly                         JSONError = 20018
	MessageEditAnnouncementRateLimit         JSONError = 20022
	UnderAge                                 JSONError = 20024
	ChannelWriteRateLimit                    JSONError = 20028
	ServerWriteRateLimit                     JSONError = 20029
	BlockedWords                             JSONError = 20031
	GuildPremiumSubscriptionLevelTooLow      JSONError = 20035
	MaxGuilds                                JSONError = 30001
	MaxFriends                               JSONError = 30002
	MaxPins                                  JSONError = 30003
	MaxRecipients                            JSONError = 30004
	MaxGuildRoles                            JSONError = 30005
	MaxWebhooks                              JSONError = 30007
	MaxEmojis                                JSONError = 30008
	MaxReactions                             JSONError = 30010
	MaxGuildChannels                         JSONError = 30013
	MaxAttachments                           JSONError = 30015
	MaxInvites                               JSONError = 30016
	MaxAnimatedEmojis                        JSONError = 30018
	MaxServerMembers                         JSONError = 30019
	MaxServerCategories                      JSONError = 30030
	GuildTemplateExists                      JSONError = 30031
	MaxApplicationCommands                   JSONError = 30032
	MaxThreadParticipants                    JSONError = 30033
	ApplicationCreateRateLimit               JSONError = 30034
	MaxNonGuildMemberBans                    JSONError = 30035
	MaxBanFetches                            JSONError = 30037
	MaxUncompletedGuildScheduledEvents       JSONError = 30038
	MaxStickers                              JSONError = 30039
	MaxPruneRequests                         JSONError = 30040
	MaxWidgetUpdates                         JSONError = 30042
	MaxOldMessageEdits                       JSONError = 30046
	ForumMaxPinnedThreads                    JSONError = 30047
	MaxForumTags                             JSONError = 30048
	BitrateTooHigh                           JSONError = 30052
	Unauthorized                             JSONError = 40001
	VerificationRequired                     JSONError = 40002
	DMOpenRateLimit                          JSONError = 40003
	SendMessagesDisabled                     JSONError = 40004
	RequestEntityTooLarge                    JSONError = 40005
	FeatureTemporarilyDisabled               JSONError = 40006
	UserGuildBanned                          JSONError = 40007
	ConnectionRevoked                        JSONError = 40012
	UserNotInVoice                           JSONError = 40032
	MessageAlreadyCrossposted                JSONError = 40033
	ApplicationNameExists                    JSONError = 40041
	ApplicationInteractionFailed             JSONError = 40043
	CannotSendInForum                        JSONError = 40058
	InteractionAlreadyAcked                  JSONError = 40060
	TagNameExists                            JSONError = 40061
	NoTagsAvailable                          JSONError = 40066
	TagRequired                              JSONError = 40067
	MissingAccess                            JSONError = 50001
	InvalidAccountType                       JSONError = 50002
	CannotExecuteOnDM                        JSONError = 50003
	GuildWidgetDisabled                      JSONError = 50004
	CannotEditMessageByOther                 JSONError = 50005
	CannotSendEmptyMessage                   JSONError = 50006
	CannotMessageUser                        JSONError = 50007
	CannotSendMessagesInChannelType          JSONError = 50008
	ChannelVerificationLevelBlock            JSONError = 50009
	OAuth2ApplicationNoBot                   JSONError = 50010
	OAuth2ApplicationLimitReached            JSONError = 50011
	InvalidOAuthState                        JSONError = 50012
	MissingPermissions                       JSONError = 50013
	InvalidAuthenticationToken               JSONError = 50014
	NoteTooLong                              JSONError = 50015
	InvalidBulkDeleteQuantity                JSONError = 50016
	InvalidMFALevel                          JSONError = 50017
	MessageCannotBePinnedInChannel           JSONError = 50019
	InvalidInvite                            JSONError = 50020
	CannotExecuteActionOnSystemMessage       JSONError = 50021
	CannotExecuteOnChannelType               JSONError = 50024
	InvalidOauthToken                        JSONError = 50025
	MissingOauthScope                        JSONError = 50026
	InvalidWebhookToken                      JSONError = 50027
	InvalidRole                              JSONError = 50028
	InvalidRecipient                         JSONError = 50033
	MessageTooOld                            JSONError = 50034
	InvalidFormBody                          JSONError = 50035
	InviteAcceptedToGuildNotContainingBot    JSONError = 50036
	InvalidActivityAction                    JSONError = 50039
	InvalidAPIVersion                        JSONError = 50041
	FileTooLarge                             JSONError = 50045
	InvalidFileUploaded                      JSONError = 50046
	CannotSelfRedeemGift                     JSONError = 50054
	InvalidGuild                             JSONError = 50055
	InvalidMessageType                       JSONError = 50068
	PaymentSourceRequired                    JSONError = 50070
	CannotModifySystemWebhook                JSONError = 50073
	CannotDeleteCommunityChannel             JSONError = 50074
	CannotEditMessageStickers                JSONError = 50080
	InvalidStickerSent                       JSONError = 50081
	ThreadOperationNotAllowed                JSONError = 50083
	InvalidThreadNotificationSettings        JSONError = 50084
	InvalidBeforeValue                       JSONError = 50085
	CommunityChannelTypeInvalid              JSONError = 50086
	ServerNotAvailable                       JSONError = 50095
	MonetizationRequired                     JSONError = 50097
	NotEnoughBoosts                          JSONError = 50101
	InvalidJSON                              JSONError = 50109
	OwnershipCannotBeTransferedToBot         JSONError = 50132
	ResizeFailed                             JSONError = 50138
	UploadedFileNotFound                     JSONError = 50146
	StickerSendMissingPermissions            JSONError = 50600
	MFARequired                              JSONError = 60003
	NoUsersWithTag                           JSONError = 80004
	ReactionBlocked                          JSONError = 90001
	ApplicationNotAvailable                  JSONError = 110001
	APIResourceOverloaded                    JSONError = 130000
	StageAlreadyOpen                         JSONError = 150006
	ReplyMissingReadHistory                  JSONError = 160002
	MessageThreadAlreadyCreated              JSONError = 160004
	ThreadLocked                             JSONError = 160005
	MaxActiveThreads                         JSONError = 160006
	MaxActiveAnnouncementThreads             JSONError = 160007
	InvalidLottieJSON                        JSONError = 170001
	LottieInvalid                            JSONError = 170002
	StickerFrameRateMaxExceeded              JSONError = 170003
	StickerFrameCountTooHigh                 JSONError = 170004
	LottieAnimationTooLong                   JSONError = 170005
	StickerFrameRateInvalid                  JSONError = 170006
	StickerAnimationDuractionTooLong         JSONError = 170007
	CannotUpdateFinishedEvent                JSONError = 180000
	CreateStageForEventFailed                JSONError = 180002
	MessageBlockedAutomod                    JSONError = 200000
	TitleBlockedAutomod                      JSONError = 200001
	ForumWebhookMissingFields                JSONError = 220001
	ForumWebhookInvalidFields                JSONError = 220002
	ForumWebookThreadsOnly                   JSONError = 220003
	ForumNoWebookServices                    JSONError = 220004
	MessageBlockedHarmful                    JSONError = 240000
)

func (e JSONError) Code() int {
	return int(e)
}

var JSONErrorMessages = map[JSONError]string{
	GeneralError:                             "Unknown error",
	UnknownAccount:                           "Unknown account",
	UnknownApplication:                       "Unknown application",
	UnknownChannel:                           "Unknown channel",
	UnknownGuild:                             "Unknown guild",
	UnknownIntegration:                       "Unknown integration",
	UnknownInvite:                            "Unknown invite",
	UnknownMember:                            "Unknown member",
	UnknownMessage:                           "Unknown message",
	UnknownPermissionOverwrite:               "Unknown permission overwrite",
	UnknownProvider:                          "Unknown provider",
	UnknownRole:                              "Unknown role",
	UnknownToken:                             "Unknown token",
	UnknownUser:                              "Unknown user",
	UnknownEmoji:                             "Unknown emoji",
	UnknownWebhook:                           "Unknown webhook",
	UnknownWebhookService:                    "Unknown webhook service",
	UnknownSession:                           "Unknown session",
	UnknownBan:                               "Unknown ban",
	UnknownSKU:                               "Unknown SKU",
	UnknownStoreListing:                      "Unknown Store Listing",
	UnknownEntitlement:                       "Unknown entitlement",
	UnknownBuild:                             "Unknown build",
	UnknownLobby:                             "Unknown lobby",
	UnknownBranch:                            "Unknown branch",
	UnknownStoreDirectoryLayout:              "Unknown store directory layout",
	UnknownRedistributable:                   "Unknown redistributable",
	UnknownGiftCode:                          "Unknown gift code",
	UnknownStream:                            "Unknown stream",
	UnknownPremiumServerSubscriptionCooldown: "Unknown premium server subscribe cooldown",
	UnknownGuildTemplate:                     "Unknown guild template",
	UnknownDiscoveryCategory:                 "Unknown discoverable server category",
	UnknownSticker:                           "Unknown sticker",
	UnknownInteraction:                       "Unknown interaction",
	UnknownApplicationCommand:                "Unknown application command",
	UnknownVoiceState:                        "Unknown voice state",
	UnknownApplicationCommandPermissions:     "Unknown application command permissions",
	UnknownStageInstance:                     "Unknown Stage Instance",
	UnknownGuildMemberVerificationForm:       "Unknown Guild Member Verification Form",
	UnknownGuildWelcomeScreen:                "Unknown Guild Welcome Screen",
	UnknownGuildScheduledEvent:               "Unknown Guild Scheduled Event",
	UnknownGuildScheduledEventUser:           "Unknown Guild Scheduled Event User",
	UnknownTag:                               "Unknown Tag",
	NoBots:                                   "Bots cannot use this endpoint",
	OnlyBots:                                 "Only bots can use this endpoint",
	ExplicitContentBlocked:                   "Explicit content cannot be sent to the desired recipient(s)",
	ApplicationUnauthorized:                  "You are not authorized to perform this action on this application",
	SlowModeRateLimit:                        "This action cannot be performed due to slowmode rate limit",
	AccountOwnerOnly:                         "Only the owner of this account can perform this action",
	MessageEditAnnouncementRateLimit:         "This message cannot be edited due to announcement rate limits",
	UnderAge:                                 "Under minimum age",
	ChannelWriteRateLimit:                    "The channel you are writing has hit the write rate limit",
	ServerWriteRateLimit:                     "The write action you are performing on the server has hit the write rate limit",
	BlockedWords:                             "Your Stage topic, server name, server description, or channel names contain words that are not allowed",
	GuildPremiumSubscriptionLevelTooLow:      "Guild premium subscription level too low",
	MaxGuilds:                                "Maximum number of guilds reached",
	MaxFriends:                               "Maximum number of friends reached",
	MaxPins:                                  "Maximum number of pins reached for the channel",
	MaxRecipients:                            "Maximum number of recipients reached",
	MaxGuildRoles:                            "Maximum number of guild roles reached",
	MaxWebhooks:                              "Maximum number of webhooks reached",
	MaxEmojis:                                "Maximum number of emojis reached",
	MaxReactions:                             "Maximum number of reactions reached",
	MaxGuildChannels:                         "Maximum number of guild channels reached",
	MaxAttachments:                           "Maximum number of attachments in a message reached",
	MaxInvites:                               "Maximum number of invites reached",
	MaxAnimatedEmojis:                        "Maximum number of animated emojis reached",
	MaxServerMembers:                         "Maximum number of server members reached",
	MaxServerCategories:                      "Maximum number of server categories has been reached",
	GuildTemplateExists:                      "Guild already has a template",
	MaxApplicationCommands:                   "Maximum number of application commands reached",
	MaxThreadParticipants:                    "Max number of thread participants has been reached",
	ApplicationCreateRateLimit:               "Max number of daily application command creates has been reached",
	MaxNonGuildMemberBans:                    "Maximum number of bans for non-guild members have been exceeded",
	MaxBanFetches:                            "Maximum number of bans fetches has been reached",
	MaxUncompletedGuildScheduledEvents:       "Maximum number of uncompleted guild scheduled events reached",
	MaxStickers:                              "Maximum number of stickers reached",
	MaxPruneRequests:                         "Maximum number of prune requests has been reached. Try again later",
	MaxWidgetUpdates:                         "Maximum number of guild widget settings updates has been reached. Try again later",
	MaxOldMessageEdits:                       "Maximum number of edits to messages older than 1 hour reached. Try again later",
	ForumMaxPinnedThreads:                    "Maximum number of pinned threads in a forum channel has been reached",
	MaxForumTags:                             "Maximum number of tags in a forum channel has been reached",
	BitrateTooHigh:                           "Bitrate is too high for channel of this type",
	Unauthorized:                             "Unauthorized. Provide a valid token and try again",
	VerificationRequired:                     "You need to verify your account in order to perform this action",
	DMOpenRateLimit:                          "You are opening direct messages too fast",
	SendMessagesDisabled:                     "Send messages has been temporarily disabled",
	RequestEntityTooLarge:                    "Request entity too large. Try sending something smaller in size",
	FeatureTemporarilyDisabled:               "This feature has been temporarily disabled server-side",
	UserGuildBanned:                          "The user is banned from this guild",
	ConnectionRevoked:                        "Connection has been revoked",
	UserNotInVoice:                           "Target user is not connected to voice",
	MessageAlreadyCrossposted:                "This message has already been crossposted",
	ApplicationNameExists:                    "An application command with that name already exists",
	ApplicationInteractionFailed:             "Application interaction failed to send",
	CannotSendInForum:                        "Cannot send a message in a forum channel",
	InteractionAlreadyAcked:                  "Interaction has already been acknowledged",
	TagNameExists:                            "Tag names must be unique",
	NoTagsAvailable:                          "There are no tags available that can be set by non-moderators",
	TagRequired:                              "A tag is required to create a forum post in this channel",
	MissingAccess:                            "Missing access",
	InvalidAccountType:                       "Invalid account type",
	CannotExecuteOnDM:                        "Cannot execute action on a DM channel",
	GuildWidgetDisabled:                      "Guild widget disabled",
	CannotEditMessageByOther:                 "Cannot edit a message authored by another user",
	CannotSendEmptyMessage:                   "Cannot send an empty message",
	CannotMessageUser:                        "Cannot send messages to this user",
	CannotSendMessagesInChannelType:          "Cannot send messages in a non-text channel",
	ChannelVerificationLevelBlock:            "Channel verification level is too high for you to gain access",
	OAuth2ApplicationNoBot:                   "OAuth2 application does not have a bot",
	OAuth2ApplicationLimitReached:            "OAuth2 application limit reached",
	InvalidOAuthState:                        "Invalid OAuth2 state",
	MissingPermissions:                       "You lack permissions to perform that action",
	InvalidAuthenticationToken:               "Invalid authentication token provided",
	NoteTooLong:                              "Note was too long",
	InvalidBulkDeleteQuantity:                "Provided too few or too many messages to delete. Must provide at least 2 and fewer than 100 messages to delete",
	InvalidMFALevel:                          "Invalid MFA Level",
	MessageCannotBePinnedInChannel:           "A message can only be pinned to the channel it was sent in",
	InvalidInvite:                            "Invite code was either invalid or taken",
	CannotExecuteActionOnSystemMessage:       "Cannot execute action on a system message",
	CannotExecuteOnChannelType:               "Cannot execute action on this channel type",
	InvalidOauthToken:                        "Invalid OAuth2 access token provided",
	MissingOauthScope:                        "Missing required OAuth2 scope",
	InvalidWebhookToken:                      "Invalid webhook token provided",
	InvalidRole:                              "Invalid role",
	InvalidRecipient:                         "Invalid Recipient(s)",
	MessageTooOld:                            "A message provided was too old to bulk delete",
	InvalidFormBody:                          "Invalid form body",
	InviteAcceptedToGuildNotContainingBot:    "An invite was accepted to a guild the application's bot is not in",
	InvalidActivityAction:                    "Invalid Activity Action",
	InvalidAPIVersion:                        "Invalid API version provided",
	FileTooLarge:                             "File uploaded exceeds the maximum size",
	InvalidFileUploaded:                      "Invalid file uploaded",
	CannotSelfRedeemGift:                     "Cannot self-redeem this gift",
	InvalidGuild:                             "Invalid Guild",
	InvalidMessageType:                       "Invalid message type",
	PaymentSourceRequired:                    "Payment source required to redeem gift",
	CannotModifySystemWebhook:                "Cannot modify a system webhook",
	CannotDeleteCommunityChannel:             "Cannot delete a channel required for Community guilds",
	CannotEditMessageStickers:                "Cannot edit stickers within a message",
	InvalidStickerSent:                       "Invalid sticker sent",
	ThreadOperationNotAllowed:                "Tried to perform an operation on an archived thread, such as editing a message or adding a user to the thread",
	InvalidThreadNotificationSettings:        "Invalid thread notification settings",
	InvalidBeforeValue:                       "before value is earlier than the thread creation date",
	CommunityChannelTypeInvalid:              "Community server channels must be text channels",
	ServerNotAvailable:                       "This server is not available in your location",
	MonetizationRequired:                     "This server needs monetization enabled in order to perform this action",
	NotEnoughBoosts:                          "This server needs more boosts to perform this action",
	InvalidJSON:                              "The request body contains invalid JSON.",
	OwnershipCannotBeTransferedToBot:         "Ownership cannot be transferred to a bot user",
	ResizeFailed:                             "Failed to resize asset below the maximum size: 262144",
	UploadedFileNotFound:                     "Uploaded file not found.",
	StickerSendMissingPermissions:            "You do not have permission to send this sticker.",
	MFARequired:                              "Two factor is required for this operation",
	NoUsersWithTag:                           "No users with DiscordTag exist",
	ReactionBlocked:                          "Reaction was blocked",
	ApplicationNotAvailable:                  "Application not yet available. Try again later",
	APIResourceOverloaded:                    "API resource is currently overloaded. Try again a little later",
	StageAlreadyOpen:                         "The Stage is already open",
	ReplyMissingReadHistory:                  "Cannot reply without permission to read message history",
	MessageThreadAlreadyCreated:              "A thread has already been created for this message",
	ThreadLocked:                             "Thread is locked",
	MaxActiveThreads:                         "Maximum number of active threads reached",
	MaxActiveAnnouncementThreads:             "Maximum number of active announcement threads reached",
	InvalidLottieJSON:                        "Invalid JSON for uploaded Lottie file",
	LottieInvalid:                            "Uploaded Lotties cannot contain rasterized images such as PNG or JPEG",
	StickerFrameRateMaxExceeded:              "Sticker maximum framerate exceeded",
	StickerFrameCountTooHigh:                 "Sticker frame count exceeds maximum of 1000 frames",
	LottieAnimationTooLong:                   "Lottie animation maximum dimensions exceeded",
	StickerFrameRateInvalid:                  "Sticker frame rate is either too small or too large",
	StickerAnimationDuractionTooLong:         "Sticker animation duration exceeds maximum of 5 seconds",
	CannotUpdateFinishedEvent:                "Cannot update a finished event",
	CreateStageForEventFailed:                "Failed to create stage needed for stage event",
	MessageBlockedAutomod:                    "Message was blocked by automatic moderation",
	TitleBlockedAutomod:                      "Title was blocked by automatic moderation",
	ForumWebhookMissingFields:                "Webhooks posted to forum channels must have a thread_name or thread_id",
	ForumWebhookInvalidFields:                "Webhooks posted to forum channels cannot have both a thread_name and thread_id",
	ForumWebookThreadsOnly:                   "Webhooks can only create threads in forum channels",
	ForumNoWebookServices:                    "Webhook services cannot be used in forum channels",
	MessageBlockedHarmful:                    "Message blocked by harmful links filter",
}
