package v1

import (
	"basefas.com/service-gin/cmd/app/handlers"
	"basefas.com/service-gin/internal/user"
	"github.com/gin-gonic/gin"
)

func LogIn(c *gin.Context) {
	var u user.Login
	if err := c.ShouldBindJSON(&u); err != nil {
		handlers.Re(c, -1, err.Error(), nil)
		return
	}
	token, err := user.Token(u)
	if err != nil {
		handlers.Re(c, -1, err.Error(), nil)
	} else {
		t := map[string]string{"token": token}
		handlers.Re(c, 0, "success", t)
	}
}
