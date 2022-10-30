package main

import (
	userscontroller "github.com/Puyodead1/fosscord-server-go/controllers"
	"github.com/Puyodead1/fosscord-server-go/initializers"
	"github.com/gin-gonic/gin"
)

func Init() {
	// Connect to database
	initializers.InitDatabase()
	initializers.InitSnowflake()
}

func main() {
	Init()
	r := gin.Default()

	r.POST("/api/v9/auth/register", userscontroller.Register)

	r.Run(":3000")
}
