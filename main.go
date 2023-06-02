package main

import (
	"io"
	"net/http"
	"os"
	"web-go-gin/controllers"
	"web-go-gin/middlewares"
	"web-go-gin/services"

	"github.com/gin-gonic/gin"
)

var (
	videoSvc services.VideoService       = services.NewVideoService()
	videoCtl controllers.VideoController = controllers.NewVideoController(videoSvc)

	loginSvc services.LoginService       = services.NewLoginService()
	jwtSvc   services.JWTService         = services.NewJWTService()
	loginCtl controllers.LoginController = controllers.NewLoginController(loginSvc, jwtSvc)
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

	// The "/view" endpoints are public (no Authorization required)
	viewRoutes := r.Group("/view")
	{
		viewRoutes.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
				"result":  http.StatusOK,
			})
		})
	}
	// JWT Authorization Middleware applies to "/api" only.
	// apiRoutes := r.Group("/api").Use(middlewares.AuthorizeJWT())
	// {
	// 	apiRoutes.GET("/videos", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, videoCtl.GetVideos())
	// 	})
	// }
	apiRoutes := r.Group("/api")
	{
		apiRoutes.POST("/login", func(ctx *gin.Context) {
			token := loginCtl.Login(ctx)
			if token != "" {
				ctx.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, nil)
			}
		})

		apiRoutes.GET("/videos", middlewares.AuthorizeJWT(), func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoCtl.GetVideos())
		})

		apiRoutes.POST("/save", middlewares.AuthorizeJWT(), func(ctx *gin.Context) {
			err := videoCtl.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is valid"})
			}
		})
	}

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

func LoadConfig() {
	panic("unimplemented")
}
