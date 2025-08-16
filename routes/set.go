package routes

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	setGroup := r.Group("/set")

	setGroup.GET("/landscape", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.File("./static/views/set/L-set.html")
	})

	setGroup.GET("/portrait", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.File("./static/views/set/P-set.html")
	})
}
