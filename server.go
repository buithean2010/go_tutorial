package main

import (
	"net/http"
	"web-go-gin/controllers"
	"web-go-gin/services"

	"github.com/gin-gonic/gin"
)

var (
	videoSvc services.VideoSvc           = services.NewVideoSvc()
	videoCtl controllers.VideoController = controllers.NewVideoCtl(videoSvc)
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"result":  http.StatusOK,
		})
	})

	r.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoCtl.GetVideos())
	})

	r.POST("/save", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoCtl.Save(ctx))
	})

	r.Run()
}
