package routes

import (
	"github.com/gin-gonic/gin"
)

func ModeRoutes(r *gin.Engine) {
	r.GET("/mode", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.File("./static/views/set/mode.html")
	})
}
