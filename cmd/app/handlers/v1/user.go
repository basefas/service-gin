package v1

import (
	"basefas.com/service-gin/internal/user"
	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	var cu user.CreateUser
	if err := c.ShouldBindJSON(&cu); err != nil {
		Re(c, -1, err.Error(), nil)
		return
	}

	err := user.Create(cu)
	if err != nil {
		Re(c, -1, err.Error(), nil)
	} else {
		Re(c, 0, "success", nil)
	}
}

func UserGet(c *gin.Context) {
	uid := c.Param("id")
	u, err := user.Get(uid)
	if err != nil {
		Re(c, -1, err.Error(), nil)
	} else {
		Re(c, 0, "success", u)
	}
}

func UserUpdate(c *gin.Context) {
	uid := c.Param("id")
	var uu user.UpdateUser
	if err := c.ShouldBindJSON(&uu); err != nil {
		Re(c, -1, err.Error(), nil)
		return
	}
	err := user.Update(uid, uu)
	if err != nil {
		Re(c, -1, err.Error(), nil)
	} else {
		Re(c, 0, "success", nil)
	}
}

func UserDelete(c *gin.Context) {
	uid := c.Param("id")
	err := user.Delete(uid)
	if err != nil {
		Re(c, -1, err.Error(), nil)
	} else {
		Re(c, 0, "success", nil)
	}
}

func UserList(c *gin.Context) {
	users, err := user.List()
	if err != nil {
		Re(c, -1, err.Error(), nil)
	} else {
		Re(c, 0, "success", users)
	}
}
