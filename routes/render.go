package routes

import "github.com/gin-gonic/gin"

func RenderRoutes(r *gin.Engine) {
	renderGroup := r.Group("/render")

	renderGroup.GET("/:subPath", func(c *gin.Context) {
		subPath := c.Param("subPath")
		c.File("./static/views/render/" + subPath + ".html")
	})
}
