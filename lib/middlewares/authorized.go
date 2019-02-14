package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Authorized blocks unauthorized requestrs
func Authorized(c *gin.Context) {
	_, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
