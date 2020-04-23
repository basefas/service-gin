package v1

import "github.com/gin-gonic/gin"

func Health(c *gin.Context) {
	Re(c, 0, "success", nil)
}
