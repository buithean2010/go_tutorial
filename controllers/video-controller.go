package controllers

import (
	"fmt"
	"web-go-gin/entities"
	"web-go-gin/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type VideoController interface {
	GetVideos() []entities.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service services.VideoSvc
}

func New(srv services.VideoSvc) *controller {
	validate = validator.New()

	return &controller{
		service: srv,
	}
}

func (c *controller) GetVideos() []entities.Video {
	return c.service.GetVideos()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entities.Video

	// Binding errors
	err := ctx.ShouldBindJSON(&video)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }

	// Validator errors
	err = validate.Struct(video)
	if err != nil {
		fmt.Println(err.Error())
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e.Namespace())
			fmt.Println(e.Field())
			fmt.Println(e.StructNamespace())
			fmt.Println(e.StructField())
			fmt.Println(e.Tag())
			fmt.Println(e.ActualTag())
			fmt.Println(e.Kind())
			fmt.Println(e.Type())
			fmt.Println(e.Value())
			fmt.Println(e.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		// https://stackoverflow.com/a/70072158
		return err
	}
	c.service.Save(video)

	return nil
}
