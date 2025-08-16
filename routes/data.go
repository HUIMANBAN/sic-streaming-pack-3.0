package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ssp/utils"
)

func DataRoutes(r *gin.Engine) {
	dataGroup := r.Group("/data")
	dataGroup.POST("/loadingData", func(c *gin.Context) {
		utils.WriteJSON[utils.LoadingData](c, "./static/data/loading.json")
	})

	dataGroup.POST("/pauseData", func(c *gin.Context) {
		utils.WriteJSON[utils.PauseData](c, "./static/data/pause.json")

	})

	dataGroup.POST("/upload", func(c *gin.Context) {
		file, formErr := c.FormFile("poster")
		if formErr != nil {
			return
		}

		saveErr := c.SaveUploadedFile(file, "./static/upload/poster")
		if saveErr != nil {
			return
		}
		fmt.Println("上传成功")

	})

	dataGroup.GET("/loadingData", func(c *gin.Context) {
		utils.ReadJSON[utils.LoadingData](c, "./static/data/loading.json")
	})

	dataGroup.GET("/pauseData", func(c *gin.Context) {
		utils.ReadJSON[utils.PauseData](c, "./static/data/pause.json")
	})
}
