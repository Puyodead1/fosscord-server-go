package gateway

import (
	"fmt"
	"log"
	"reflect"

	"github.com/Puyodead1/fosscord-server-go/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kataras/go-events"
)

var Events = events.New()

type Client struct {
	ID           string
	Conn         *websocket.Conn
	Pool         *Pool
	Events       map[string]func()
	MemberEvents map[string]func()
	Sequence     int
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

type Event struct {
	GuildID   *string
	UserID    *string
	ChannelID *string
	CreatedAt *string
	EventName EVENT
	Data      any
}

type EventOpts struct {
	Event
	Acknowledge *func()
	Cancel      func()
}
type GuildMemberAdd struct {
	models.Member
	GuildID string `json:"guild_id"`
}

type GuildCreate struct {
	models.Guild
	JoinedAt string `json:"joined_at"` // when the current user joined the guild
	// TODO: scheduled events
	// TODO: guild hashes
	// TODO: presences
	// TODO: stage instances
	// TODO: threads
	// TODO: embeded activities
}

var ClientPool Pool

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		payload := IncomingMessage{Client: c, Message: message}
		c.Pool.Message <- payload
		log.Printf("Message Received: %+v\n", message)
	}
}

func serve(pool *Pool, c *gin.Context) {
	log.Println("WebSocket Endpoint Hit")
	conn, err := Upgrade(c.Writer, c.Request)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	client := &Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func Init() {
	log.Println("Starting Gateway")

	r := gin.Default()

	ClientPool = *NewPool()
	go ClientPool.Start()

	r.GET("/", func(ctx *gin.Context) {
		serve(&ClientPool, ctx)
	})

	r.Run(":3001")
}

type EVENT string

const (
	EVENT_READY                         EVENT = "READY"
	EVENT_CHANNEL_CREATE                EVENT = "CHANNEL_CREATE"
	EVENT_CHANNEL_UPDATE                EVENT = "CHANNEL_UPDATE"
	EVENT_CHANNEL_DELETE                EVENT = "CHANNEL_DELETE"
	EVENT_CHANNEL_PINS_UPDATE           EVENT = "CHANNEL_PINS_UPDATE"
	EVENT_CHANNEL_RECIPIENT_ADD         EVENT = "CHANNEL_RECIPIENT_ADD"
	EVENT_CHANNEL_RECIPIENT_REMOVE      EVENT = "CHANNEL_RECIPIENT_REMOVE"
	EVENT_GUILD_CREATE                  EVENT = "GUILD_CREATE"
	EVENT_GUILD_UPDATE                  EVENT = "GUILD_UPDATE"
	EVENT_GUILD_DELETE                  EVENT = "GUILD_DELETE"
	EVENT_GUILD_BAN_ADD                 EVENT = "GUILD_BAN_ADD"
	EVENT_GUILD_BAN_REMOVE              EVENT = "GUILD_BAN_REMOVE"
	EVENT_GUILD_EMOJIS_UPDATE           EVENT = "GUILD_EMOJIS_UPDATE"
	EVENT_GUILD_STICKERS_UPDATE         EVENT = "GUILD_STICKERS_UPDATE"
	EVENT_GUILD_INTEGRATIONS_UPDATE     EVENT = "GUILD_INTEGRATIONS_UPDATE"
	EVENT_GUILD_MEMBER_ADD              EVENT = "GUILD_MEMBER_ADD"
	EVENT_GUILD_MEMBER_REMOVE           EVENT = "GUILD_MEMBER_REMOVE"
	EVENT_GUILD_MEMBER_UPDATE           EVENT = "GUILD_MEMBER_UPDATE"
	EVENT_GUILD_MEMBER_SPEAKING         EVENT = "GUILD_MEMBER_SPEAKING"
	EVENT_GUILD_MEMBERS_CHUNK           EVENT = "GUILD_MEMBERS_CHUNK"
	EVENT_GUILD_MEMBER_LIST_UPDATE      EVENT = "GUILD_MEMBER_LIST_UPDATE"
	EVENT_GUILD_ROLE_CREATE             EVENT = "GUILD_ROLE_CREATE"
	EVENT_GUILD_ROLE_UPDATE             EVENT = "GUILD_ROLE_UPDATE"
	EVENT_GUILD_ROLE_DELETE             EVENT = "GUILD_ROLE_DELETE"
	EVENT_INVITE_CREATE                 EVENT = "INVITE_CREATE"
	EVENT_INVITE_DELETE                 EVENT = "INVITE_DELETE"
	EVENT_MESSAGE_CREATE                EVENT = "MESSAGE_CREATE"
	EVENT_MESSAGE_UPDATE                EVENT = "MESSAGE_UPDATE"
	EVENT_MESSAGE_DELETE                EVENT = "MESSAGE_DELETE"
	EVENT_MESSAGE_DELETE_BULK           EVENT = "MESSAGE_DELETE_BULK"
	EVENT_MESSAGE_REACTION_ADD          EVENT = "MESSAGE_REACTION_ADD"
	EVENT_MESSAGE_REACTION_REMOVE       EVENT = "MESSAGE_REACTION_REMOVE"
	EVENT_MESSAGE_REACTION_REMOVE_ALL   EVENT = "MESSAGE_REACTION_REMOVE_ALL"
	EVENT_MESSAGE_REACTION_REMOVE_EMOJI EVENT = "MESSAGE_REACTION_REMOVE_EMOJI"
	EVENT_PRESENCE_UPDATE               EVENT = "PRESENCE_UPDATE"
	EVENT_TYPING_START                  EVENT = "TYPING_START"
	EVENT_USER_UPDATE                   EVENT = "USER_UPDATE"
	EVENT_USER_NOTE_UPDATE              EVENT = "USER_NOTE_UPDATE"
	EVENT_WEBHOOKS_UPDATE               EVENT = "WEBHOOKS_UPDATE"
	EVENT_INTERACTION_CREATE            EVENT = "INTERACTION_CREATE"
	EVENT_VOICE_STATE_UPDATE            EVENT = "VOICE_STATE_UPDATE"
	EVENT_VOICE_SERVER_UPDATE           EVENT = "VOICE_SERVER_UPDATE"
	EVENT_APPLICATION_COMMAND_CREATE    EVENT = "APPLICATION_COMMAND_CREATE"
	EVENT_APPLICATION_COMMAND_UPDATE    EVENT = "APPLICATION_COMMAND_UPDATE"
	EVENT_APPLICATION_COMMAND_DELETE    EVENT = "APPLICATION_COMMAND_DELETE"
	EVENT_MESSAGE_ACK                   EVENT = "MESSAGE_ACK"
	EVENT_RELATIONSHIP_ADD              EVENT = "RELATIONSHIP_ADD"
	EVENT_RELATIONSHIP_REMOVE           EVENT = "RELATIONSHIP_REMOVE"
	EVENT_SESSIONS_REPLACE              EVENT = "SESSIONS_REPLACE"
)

func (e EVENT) String() string {
	return string(e)
}

func ToEvent(s string) EVENT {
	switch s {
	case "READY":
		return EVENT_READY
	case "CHANNEL_CREATE":
		return EVENT_CHANNEL_CREATE
	case "CHANNEL_UPDATE":
		return EVENT_CHANNEL_UPDATE
	case "CHANNEL_DELETE":
		return EVENT_CHANNEL_DELETE
	case "CHANNEL_PINS_UPDATE":
		return EVENT_CHANNEL_PINS_UPDATE
	case "CHANNEL_RECIPIENT_ADD":
		return EVENT_CHANNEL_RECIPIENT_ADD
	case "CHANNEL_RECIPIENT_REMOVE":
		return EVENT_CHANNEL_RECIPIENT_REMOVE
	case "GUILD_CREATE":
		return EVENT_GUILD_CREATE
	case "GUILD_UPDATE":
		return EVENT_GUILD_UPDATE
	case "GUILD_DELETE":
		return EVENT_GUILD_DELETE
	case "GUILD_BAN_ADD":
		return EVENT_GUILD_BAN_ADD
	case "GUILD_BAN_REMOVE":
		return EVENT_GUILD_BAN_REMOVE
	case "GUILD_EMOJIS_UPDATE":
		return EVENT_GUILD_EMOJIS_UPDATE
	case "GUILD_STICKERS_UPDATE":
		return EVENT_GUILD_STICKERS_UPDATE
	case "GUILD_INTEGRATIONS_UPDATE":
		return EVENT_GUILD_INTEGRATIONS_UPDATE
	case "GUILD_MEMBER_ADD":
		return EVENT_GUILD_MEMBER_ADD
	case "GUILD_MEMBER_REMOVE":
		return EVENT_GUILD_MEMBER_REMOVE
	case "GUILD_MEMBER_UPDATE":
		return EVENT_GUILD_MEMBER_UPDATE
	case "GUILD_MEMBER_SPEAKING":
		return EVENT_GUILD_MEMBER_SPEAKING
	case "GUILD_MEMBERS_CHUNK":
		return EVENT_GUILD_MEMBERS_CHUNK
	case "GUILD_MEMBER_LIST_UPDATE":
		return EVENT_GUILD_MEMBER_LIST_UPDATE
	case "GUILD_ROLE_CREATE":
		return EVENT_GUILD_ROLE_CREATE
	case "GUILD_ROLE_UPDATE":
		return EVENT_GUILD_ROLE_UPDATE
	case "GUILD_ROLE_DELETE":
		return EVENT_GUILD_ROLE_DELETE
	case "INVITE_CREATE":
		return EVENT_INVITE_CREATE
	case "INVITE_DELETE":
		return EVENT_INVITE_DELETE
	case "MESSAGE_CREATE":
		return EVENT_MESSAGE_CREATE
	case "MESSAGE_UPDATE":
		return EVENT_MESSAGE_UPDATE
	case "MESSAGE_DELETE":
		return EVENT_MESSAGE_DELETE
	case "MESSAGE_DELETE_BULK":
		return EVENT_MESSAGE_DELETE_BULK
	case "MESSAGE_REACTION_ADD":
		return EVENT_MESSAGE_REACTION_ADD
	case "MESSAGE_REACTION_REMOVE":
		return EVENT_MESSAGE_REACTION_REMOVE
	case "MESSAGE_REACTION_REMOVE_ALL":
		return EVENT_MESSAGE_REACTION_REMOVE_ALL
	case "MESSAGE_REACTION_REMOVE_EMOJI":
		return EVENT_MESSAGE_REACTION_REMOVE_EMOJI
	case "PRESENCE_UPDATE":
		return EVENT_PRESENCE_UPDATE
	case "TYPING_START":
		return EVENT_TYPING_START
	case "USER_UPDATE":
		return EVENT_USER_UPDATE
	case "USER_NOTE_UPDATE":
		return EVENT_USER_NOTE_UPDATE
	case "WEBHOOKS_UPDATE":
		return EVENT_WEBHOOKS_UPDATE
	case "INTERACTION_CREATE":
		return EVENT_INTERACTION_CREATE
	case "VOICE_STATE_UPDATE":
		return EVENT_VOICE_STATE_UPDATE
	case "VOICE_SERVER_UPDATE":
		return EVENT_VOICE_SERVER_UPDATE
	case "APPLICATION_COMMAND_CREATE":
		return EVENT_APPLICATION_COMMAND_CREATE
	case "APPLICATION_COMMAND_UPDATE":
		return EVENT_APPLICATION_COMMAND_UPDATE
	case "APPLICATION_COMMAND_DELETE":
		return EVENT_APPLICATION_COMMAND_DELETE
	case "MESSAGE_ACK":
		return EVENT_MESSAGE_ACK
	case "RELATIONSHIP_ADD":
		return EVENT_RELATIONSHIP_ADD
	case "RELATIONSHIP_REMOVE":
		return EVENT_RELATIONSHIP_REMOVE
	case "SESSIONS_REPLACE":
		return EVENT_SESSIONS_REPLACE
	}
	return ""
}

func EmitEvent(payload Event) error {
	var id *string

	if payload.ChannelID != nil {
		id = payload.ChannelID
	} else if payload.UserID != nil {
		id = payload.UserID
	} else {
		id = payload.GuildID
	}

	if id == nil {
		// create an error
		return fmt.Errorf("event doesnt contain any id %v", payload)
	}

	log.Printf("Emitting event %s for %s", payload.EventName, *id)

	Events.Emit(events.EventName(*id), payload)
	return nil
}

func handlePresenceUpdate(args ...interface{}) {
	// TODO:
	a := args[0]
	log.Printf("handlePresenceUpdate %v", a)
	log.Printf("Type: %v", reflect.TypeOf(a))
}

func consume(args ...interface{}) {
	event := args[0].(Event)
	client := args[1].(*Client)

	var id string
	if event.ChannelID != nil {
		id = *event.ChannelID
	} else if event.UserID != nil {
		id = *event.UserID
	} else {
		id = *event.GuildID
	}
	log.Printf("Consuming event %s for %s", event.EventName, id)

	// TODO: permissions check
	switch event.EventName {
	case EVENT_GUILD_MEMBER_REMOVE:
		e := client.MemberEvents[id]
		if e != nil {
			e()
		}
		delete(client.MemberEvents, id)
	case EVENT_GUILD_MEMBER_ADD:
		e := client.MemberEvents[id]
		if e != nil {
			break // already subscribed
		}
		// append to the list of events
		client.MemberEvents[id] = ListenEvent(id, client, handlePresenceUpdate)
	case EVENT_GUILD_MEMBER_UPDATE:
		e := client.MemberEvents[id]
		if e != nil {
			e()
		}
	case EVENT_RELATIONSHIP_REMOVE:
	case EVENT_CHANNEL_DELETE:
	case EVENT_GUILD_DELETE:
		delete(client.Events, id)
		// TODO: opts.cancel()
	case EVENT_CHANNEL_CREATE:
		// TODO: permissions check
		client.Events[id] = ListenEvent(id, client, consume)
	case EVENT_RELATIONSHIP_ADD:
		client.Events[id] = ListenEvent(id, client, handlePresenceUpdate)
	case EVENT_GUILD_CREATE:
		log.Printf("EVENT_GUILD_CREATE %v", event)
		client.Events[id] = ListenEvent(id, client, consume)
	case EVENT_CHANNEL_UPDATE:
		e := client.Events[id]
		// TODO: permissions check
		if e == nil {
			return
		}
		// TODO: opts.cancel()
		delete(client.Events, id)
	}

	// TODO: permission checking

	log.Printf("Emitting event %s for %s", event.EventName, id)

	// TODO: sequence
	payload := GatewayPayload{
		Op:  GATEWAYOPCODE_DISPATCH.Value(),
		T:   event.EventName.String(),
		D:   event.Data,
		Seq: client.Sequence,
	}

	log.Printf("Sending payload %v", payload)

	client.Conn.WriteJSON(payload)
}

func ListenEvent(id string, c *Client, callback func(...interface{})) func() {
	// check if we already have a listener for this event
	e := c.Events[id]
	if e != nil {
		return e
	}
	log.Printf("Adding listener for %s", id)
	listener := func(args ...interface{}) {
		log.Printf("Listener for %s called", id)
		// add the client to the args
		args = append(args, c)
		callback(args...)
	}
	log.Printf("creating cancel callback for %s", id)
	cancel := func() {
		log.Printf("Removing listener for %s", id)
		Events.RemoveListener(events.EventName(id), listener)
		Events.SetMaxListeners(Events.GetMaxListeners() - 1)
	}
	Events.SetMaxListeners(Events.GetMaxListeners() + 1)
	Events.AddListener(events.EventName(id), listener)

	log.Printf("returning cancel for %s", id)
	return cancel
}

func SetupListeners(id string, c *Client) {
	log.Printf("Setting up listeners for %s", id)
	c.Events[id] = ListenEvent(id, c, consume)

	// TODO: relationships
	// TODO: dms
	// TODO: guilds

}
