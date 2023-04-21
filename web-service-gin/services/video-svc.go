package services

import "web-go-gin/entities"

type VideoSvc interface {
	GetVideos() []entities.Video
	Save(entities.Video) entities.Video
}

type videoSvc struct {
	videos []entities.Video
}

func NewVideoSvc() *videoSvc {
	return &videoSvc{
		videos: []entities.Video{},
	}
}

func (svc *videoSvc) GetVideos() []entities.Video {
	return svc.videos
}

func (svc *videoSvc) Save(vid entities.Video) entities.Video {
	svc.videos = append(svc.videos, vid)
	return vid
}
