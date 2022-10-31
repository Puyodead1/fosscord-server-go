package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	authcontroller "github.com/Puyodead1/fosscord-server-go/controllers/auth"
	userscontroller "github.com/Puyodead1/fosscord-server-go/controllers/users"
	"github.com/Puyodead1/fosscord-server-go/gateway"
	"github.com/Puyodead1/fosscord-server-go/initializers"
	"github.com/Puyodead1/fosscord-server-go/middleware"
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
		// serve cached files if they exist
		filename := path.Base(c.Request.URL.String())
		filename = strings.Split(filename, "?")[0] // remove query string

		if _, err := os.Stat("./client/cache/" + filename); !os.IsNotExist(err) {
			c.File("./client/cache/" + filename)
			return
		}

		resp, err := http.Get("https://canary.discord.com" + c.Request.URL.Path)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = os.MkdirAll("./client/cache", 0644)
		if err != nil {
			log.Printf("Error creating cache folder: %s", err.Error())
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// write to cache file
		err = os.WriteFile("./client/cache/"+filename, body, 0644)
		if err != nil {
			log.Printf("Error writing file: %s", err.Error())
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.Data(200, resp.Header.Get("Content-Type"), body)
	})

	api := r.Group("/api/v9")

	api.POST("auth/register", authcontroller.Register)
	api.POST("auth/login", authcontroller.Login)

	apiProtected := r.Group("/api/v9")
	apiProtected.Use(middleware.JwtAuthMiddleware())
	apiProtected.GET("/users/:id/affinities/guilds", userscontroller.GetGuildAffinities)
	apiProtected.GET("/users/:id/affinities/users", userscontroller.GetUserAffinities)
	apiProtected.GET("/users/:id/library", userscontroller.GetLibrary)
	apiProtected.GET("/users/:id/billing/localized-pricing-promo", userscontroller.GetBillingLocalizedPricingPromo)
	apiProtected.GET("/users/:id/billing/payment-sources", userscontroller.GetBillingPaymentSources)
	apiProtected.GET("/users/:id/billing/country-code", userscontroller.GetBillingCountryCode)

	r.Run(":3000")
}

func PatchAssets() {
	log.Println("Patching assets")
	// loop all .js files in ./client/cache
	files, err := os.ReadDir("./client/cache")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".js") {
			continue
		}

		// read file
		data, err := os.ReadFile("./client/cache/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		data = bytes.ReplaceAll(data, []byte("e.isDiscordGatewayPlaintextSet=function(){0;return!1};"), []byte("e.isDiscordGatewayPlaintextSet = function() { return true };"))
		data = bytes.ReplaceAll(data, []byte("//# sourceMappingURL="), []byte("//# disabledSourceMappingURL="))
		data = bytes.ReplaceAll(data, []byte("https://fa97a90475514c03a42f80cd36d147c4@sentry.io/140984"), []byte("https://6bad92b0175d41a18a037a73d0cff282@sentry.thearcanebrony.net/12"))

		// write file
		err = os.WriteFile("./client/cache/"+file.Name(), data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	Init()
	// patch cached assets
	PatchAssets()
	// start the servers
	go StartAPI()
	gateway.Init()
	// TODO: cdn
}
