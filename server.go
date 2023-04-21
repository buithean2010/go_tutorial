package main

import (
	"io"
	"net/http"
	"os"
	"web-go-gin/controllers"
	"web-go-gin/services"

	"github.com/gin-gonic/gin"
)

var (
	videoSvc services.VideoSvc           = services.NewVideoSvc()
	videoCtl controllers.VideoController = controllers.NewVideoCtl(videoSvc)
)

func setupLogger(toFile bool) {
	if toFile {
		// Disable Console Color, you don't need console color when writing the logs to file.
		gin.DisableConsoleColor()

		// Logging to a file.
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)

		// Logging to a file AND console at the same time.
		//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
}

func main() {
	setupLogger(false)

	r := gin.New()

	r.Use(gin.Recovery())

	// r.Use(middlewares.Logger())

	// r.Use(gindump.Dump())

	// r.Use(middlewares.BasicAuth())

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
