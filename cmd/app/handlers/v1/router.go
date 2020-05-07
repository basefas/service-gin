package v1

import (
	"fmt"

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

	r.GET("/health", Health)

	api := r.Group("/api/v1")

	user := api.Group("/user")
	user.Use(middleware.JWT())
	{
		user.POST("", UserCreate)
		user.GET("/:id", UserGet)
		user.PUT("/:id", UserUpdate)
		user.DELETE("/:id", UserDelete)
		user.GET("/", UserList)
	}

	policy := api.Group("/policy")
	policy.Use(middleware.Casbin(auth.Casbin))

	{
		policy.POST("", PolicyCreate)
		policy.GET("/:id", PolicyGet)
	}

	login := api.Group("/login")
	{
		login.POST("", LogIn)
	}

	return r
}
