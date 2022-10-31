package gateway

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
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
			packet := GatewayPayload{Op: 10, D: map[string]interface{}{"heartbeat_interval": 41250}}
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
				log.Println("IDENTIFY")
				break
			default:
				log.Printf("Unknown OP Code: %d", payload.Op)
				cm := websocket.FormatCloseMessage(UNKNOWN_OPCODE.Value(), CloseCodeMessages[UNKNOWN_OPCODE])
				if err := message.Client.Conn.WriteMessage(websocket.CloseMessage, cm); err != nil {
					log.Println(err)
				}
				message.Client.Conn.Close()
				break
			}
		}
	}
}
