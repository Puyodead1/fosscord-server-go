package gateway

import (
	"encoding/json"
	"log"

	userservices "github.com/Puyodead1/fosscord-server-go/services"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

type IncomingMessage struct {
	Client  *Client
	Message Message
}

type GatewayPayload struct {
	Op  int    `json:"op"`
	Seq int    `json:"s,omitempty"`
	T   string `json:"t,omitempty"`
	D   any    `json:"d,omitempty"`
}

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Message    chan IncomingMessage
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Message:    make(chan IncomingMessage),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Println("Size of Connection Pool: ", len(pool.Clients))
			log.Println(client)
			packet := GatewayPayload{Op: GATEWAYOPCODE_HELLO.Value(), D: map[string]interface{}{"heartbeat_interval": 41250, "_trace": []string{"[\"localhost\",{\"micros\":0.0}]"}}}
			client.Conn.WriteJSON(packet)
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Println("Size of Connection Pool: ", len(pool.Clients))
			break
		case message := <-pool.Message:
			payload := GatewayPayload{}
			// unmarshal message into payload
			err := json.Unmarshal([]byte(message.Message.Body), &payload)
			if err != nil {
				cm := websocket.FormatCloseMessage(DECODE_ERROR.Value(), CloseCodeMessages[DECODE_ERROR])
				if err := message.Client.Conn.WriteMessage(websocket.CloseMessage, cm); err != nil {
					log.Println(err)
				}
				message.Client.Conn.Close()
				break
			}

			// switch the op code
			switch payload.Op {
			case GATEWAYOPCODE_HEARTBEAT.Value():
				log.Println("Heartbeat")
				message.Client.Conn.WriteJSON(GatewayPayload{Op: GATEWAYOPCODE_HEARTBEAT_ACK.Value()})
				break
			case GATEWAYOPCODE_IDENTIFY.Value():

				identifyPayload := IdentifyPayload{}
				mapstructure.Decode(payload.D, &identifyPayload) // TODO: validate the fields

				id, err := userservices.VerifyToken(identifyPayload.Token)
				if err != nil {
					cm := websocket.FormatCloseMessage(AUTHENTICATION_FAILED.Value(), CloseCodeMessages[AUTHENTICATION_FAILED])
					if err := message.Client.Conn.WriteMessage(websocket.CloseMessage, cm); err != nil {
						log.Println(err)
					}
					message.Client.Conn.Close()
					break
				}

				user := userservices.GetUserById(id)
				if user.ID == "" {
					cm := websocket.FormatCloseMessage(AUTHENTICATION_FAILED.Value(), CloseCodeMessages[AUTHENTICATION_FAILED])
					if err := message.Client.Conn.WriteMessage(websocket.CloseMessage, cm); err != nil {
						log.Println(err)
					}
					message.Client.Conn.Close()
					break
				}

				sessionId := userservices.GenerateSessionID()
				if sessionId == "" {
					cm := websocket.FormatCloseMessage(UNKNOWN_ERROR.Value(), CloseCodeMessages[UNKNOWN_ERROR])
					if err := message.Client.Conn.WriteMessage(websocket.CloseMessage, cm); err != nil {
						log.Println(err)
					}
					message.Client.Conn.Close()
					break
				}

				userSettings := userservices.GetUserSettings(id)
				if userSettings.ID == "" {
					cm := websocket.FormatCloseMessage(UNKNOWN_ERROR.Value(), CloseCodeMessages[UNKNOWN_ERROR])
					if err := message.Client.Conn.WriteMessage(websocket.CloseMessage, cm); err != nil {
						log.Println(err)
					}
					message.Client.Conn.Close()
					break
				}

				readStates := userservices.GetReadStates(id)
				if readStates == nil {
					cm := websocket.FormatCloseMessage(UNKNOWN_ERROR.Value(), CloseCodeMessages[UNKNOWN_ERROR])
					if err := message.Client.Conn.WriteMessage(websocket.CloseMessage, cm); err != nil {
						log.Println(err)
					}
					message.Client.Conn.Close()
					break
				}

				readStateEntries := make([]any, len(readStates))
				for i, readState := range readStates {
					readStateEntries[i] = readState.ChannelID
				}

				empty := make([]any, 0)

				readyPayload := ReadyEventPayload{
					V:                9,
					User:             user,
					PrivateChannels:  empty,
					SessionID:        sessionId,
					Guilds:           empty,
					UserSettings:     &userSettings,
					Users:            &empty,
					Experiments:      &empty,
					GuildExperiments: &empty,
					ReadState: &ReadyEventDataStruct1{
						Entries: readStateEntries,
						Version: 304128,
						Partial: false,
					},
					UserGuildSettings: &ReadyEventDataStruct1{
						Entries: empty,
						Version: 642,
						Partial: false,
					},
					ConnectedAccounts: &empty,
				}
				message.Client.Conn.WriteJSON(GatewayPayload{Op: GATEWAYOPCODE_DISPATCH.Value(), T: "READY", D: readyPayload})
				break
			default:
				log.Printf("Unknown OP Code: %d", payload.Op)
				// TODO: uncomment when all opcodes are implemented
				// cm := websocket.FormatCloseMessage(UNKNOWN_OPCODE.Value(), CloseCodeMessages[UNKNOWN_OPCODE])
				// if err := message.Client.Conn.WriteMessage(websocket.CloseMessage, cm); err != nil {
				// 	log.Println(err)
				// }
				// message.Client.Conn.Close()
				break
			}
		}
	}
}
