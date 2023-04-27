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
	videoSvc services.VideoSvc           = services.New()
	videoCtl controllers.VideoController = controllers.New(videoSvc)
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
		err := videoCtl.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is valid"})
		}
	})

	r.Run()
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        r,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()
}
