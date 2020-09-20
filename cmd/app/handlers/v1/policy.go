package v1

import (
	"basefas.com/service-gin/cmd/app/handlers"
	"basefas.com/service-gin/internal/auth"
	"github.com/gin-gonic/gin"
)

func PolicyCreate(c *gin.Context) {
	var cp auth.CreatePolicy
	if err := c.ShouldBindJSON(&cp); err != nil {
		handlers.Re(c, -1, err.Error(), nil)
		return
	}

	_, err := auth.Casbin.AddPolicy(cp.UID, cp.PolicyUrl, cp.PolicyMethod)

	if err != nil {
		handlers.Re(c, -1, err.Error(), nil)
	} else {
		handlers.Re(c, 0, "success", nil)
	}
}

func PolicyGet(c *gin.Context) {
	uid := c.Param("id")

	res, err := auth.Casbin.GetRolesForUser(uid)

	if err != nil {
		handlers.Re(c, -1, err.Error(), nil)

	} else {
		handlers.Re(c, 0, "success", res)
	}
}
