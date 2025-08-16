package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
	"log"
	"os"
	"ssp/routes"
	"ssp/utils"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Static("/static", "./static")

	//首页
	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/mode")
	})

	//mode路由分组
	routes.ModeRoutes(r)

	//set路由分组
	routes.SetRoutes(r)

	//data路由分组
	routes.DataRoutes(r)

	//render路由分组
	routes.RenderRoutes(r)

	var port = "59137"

	//打开链接
	openErr := browser.OpenURL("http://localhost:" + port)
	if openErr != nil {
		log.Println(openErr)
	}

	utils.PrintInfo()
	log.Printf("Server is listening on port %s\n", port)
	fmt.Println("按住Ctrl键并点击链接即可打开 http://localhost:" + port)

	runErr := r.Run("0.0.0.0:" + port)
	if runErr != nil {
		log.Println(runErr)
		os.Exit(1)
	}
}
