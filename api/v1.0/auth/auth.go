package auth

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", register)
		auth.POST("/login", login)
		auth.GET("/check", check)
	}
}
