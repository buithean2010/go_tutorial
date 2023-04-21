package controllers

import (
	"web-go-gin/entities"
	"web-go-gin/services"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	GetVideos() []entities.Video
	Save(ctx *gin.Context) entities.Video
}

type controller struct {
	service services.VideoSvc
}

func NewVideoCtl(srv services.VideoSvc) *controller {
	return &controller{
		service: srv,
	}
}

func (c *controller) GetVideos() []entities.Video {
	return c.service.GetVideos()
}

func (c *controller) Save(ctx *gin.Context) entities.Video {
	var video entities.Video
	ctx.BindJSON(&video)
	c.service.Save(video)

	return video
}
