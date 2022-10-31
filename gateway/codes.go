package gateway

type OPCode int

const (
	GATEWAYOPCODE_DISPATCH OPCode = iota
	GATEWAYOPCODE_HEARTBEAT
	GATEWAYOPCODE_IDENTIFY
	GATEWAYOPCODE_STATUS_UPDATE
	GATEWAYOPCODE_VOICE_STATE_UPDATE
	GATEWAYOPCODE_VOICE_SERVER_PING
	GATEWAYOPCODE_RESUME
	GATEWAYOPCODE_RECONNECT
	GATEWAYOPCODE_REQUEST_GUILD_MEMBERS
	GATEWAYOPCODE_INVALID_SESSION
	GATEWAYOPCODE_HELLO
	GATEWAYOPCODE_HEARTBEAT_ACK
)

func (g OPCode) Value() int {
	return int(g)
}

type CloseCode int

const (
	UNKNOWN_ERROR CloseCode = iota + 4000
	UNKNOWN_OPCODE
	DECODE_ERROR
	NOT_AUTHENTICATED
	AUTHENTICATION_FAILED
	ALREADY_AUTHENTICATED
	INVALID_SEQ
	RATE_LIMITED
	SESSION_TIMED_OUT
	INVALID_SHARD
	SHARDING_REQUIRED
	INVALID_API_VERSION
	INVALID_INTENTS
	DISALLOWED_INTENTS
)

func (c CloseCode) Value() int {
	return int(c)
}

var CloseCodeMessages = map[CloseCode]string{
	UNKNOWN_ERROR:         "An unknown error occurred",
	UNKNOWN_OPCODE:        "An unknown opcode was sent",
	DECODE_ERROR:          "Failed to decode payload",
	NOT_AUTHENTICATED:     "Sent a payload prior to identifying",
	AUTHENTICATION_FAILED: "The account token sent with your identify payload is incorrect",
	ALREADY_AUTHENTICATED: "More than one identify payload was sent",
	INVALID_SEQ:           "The sequence sent when resuming the session was invalid",
	RATE_LIMITED:          "Woah nelly! You're sending payloads to us too quickly. Slow it down! We will resume your session automatically",
	SESSION_TIMED_OUT:     "Your session timed out. Reconnect and start a new one",
	INVALID_SHARD:         "You sent us an invalid shard when identifying",
	SHARDING_REQUIRED:     "The session would have handled too many guilds - you are required to shard your connection in order to connect",
	INVALID_API_VERSION:   "You sent an invalid version for the gateway",
	INVALID_INTENTS:       "You sent an invalid intent for a Gateway Intent. You may have incorrectly calculated the bitwise value.",
	DISALLOWED_INTENTS:    "You sent a disallowed intent for a Gateway Intent. You may have tried to specify an intent that you have not enabled or are not whitelisted for.",
}
