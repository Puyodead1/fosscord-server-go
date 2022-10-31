package main

import (
	"io/ioutil"
	"log"
	"net/http"

	userscontroller "github.com/Puyodead1/fosscord-server-go/controllers"
	"github.com/Puyodead1/fosscord-server-go/gateway"
	"github.com/Puyodead1/fosscord-server-go/initializers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Init() {
	// Connect to database
	initializers.InitDatabase()
	initializers.InitSnowflake()
	binding.Validator = new(initializers.DefaultValidator)
}

func StartAPI() {
	log.Println("Starting API")
	r := gin.Default()
	r.StaticFile("/", "./client/login.html")

	// Proxies assets to discord
	r.Any("/assets/:file", func(c *gin.Context) {
		path := c.Request.URL.Path
		resp, err := http.Get("https://canary.discord.com" + path)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.Data(200, resp.Header.Get("Content-Type"), body)
	})

	api := r.Group("/api/v9")

	api.POST("auth/register", userscontroller.Register)
	api.POST("auth/login", userscontroller.Login)

	r.Run(":3000")
}

func main() {
	Init()
	// start the servers
	go StartAPI()
	gateway.Init()
	// TODO: cdn
}
