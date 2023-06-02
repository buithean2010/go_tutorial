package services

import "web-go-gin/entities"

type VideoService interface {
	GetVideos() []entities.Video
	Save(entities.Video) entities.Video
}

type videoService struct {
	videos []entities.Video
}

func NewVideoService() *videoService {
	return &videoService{
		videos: []entities.Video{},
	}
}

func (svc *videoService) GetVideos() []entities.Video {
	return svc.videos
}

func (svc *videoService) Save(vid entities.Video) entities.Video {
	svc.videos = append(svc.videos, vid)
	return vid
}
