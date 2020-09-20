package handlers

import (
	"fmt"

	"basefas.com/service-gin/cmd/app/handlers/base"
	"basefas.com/service-gin/cmd/app/handlers/v1"
	"basefas.com/service-gin/internal/auth"
	middleware "basefas.com/service-gin/internal/mid"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() {
	r := setupRouter()
	port := fmt.Sprintf(":%s", viper.GetString("app.port"))
	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}

func setMode() {
	switch viper.GetString("app.runMode") {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		fmt.Println("Load App Mode Error!")
	}
}

func setupRouter() *gin.Engine {
	setMode()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", base.Health)

	api := r.Group("/api/v1")

	user := api.Group("/user")
	user.Use(middleware.JWT())
	{
		user.POST("", v1.UserCreate)
		user.GET("/:id", v1.UserGet)
		user.PUT("/:id", v1.UserUpdate)
		user.DELETE("/:id", v1.UserDelete)
		user.GET("/", v1.UserList)
	}

	policy := api.Group("/policy")
	policy.Use(middleware.Casbin(auth.Casbin))

	{
		policy.POST("", v1.PolicyCreate)
		policy.GET("/:id", v1.PolicyGet)
	}

	login := api.Group("/login")
	{
		login.POST("", v1.LogIn)
	}

	return r
}
