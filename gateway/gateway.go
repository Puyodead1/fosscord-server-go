package gateway

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

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

	pool := NewPool()
	go pool.Start()

	r.GET("/", func(ctx *gin.Context) {
		serve(pool, ctx)
	})

	r.Run(":3001")
}
