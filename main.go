package main

import (
	userscontroller "github.com/Puyodead1/fosscord-server-go/controllers"
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

func main() {
	Init()
	r := gin.Default()

	api := r.Group("/api/v9")

	api.POST("auth/register", userscontroller.Register)
	api.POST("auth/login", userscontroller.Login)

	r.Run(":3000")
}
