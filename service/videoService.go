package service

import (
	"TikTokApp/common"
	"TikTokApp/models"
	"TikTokApp/utils"
	"mime/multipart"
)

var videoDao = models.NewVideoDaoInstance()

// Publish 发布视频
func Publish(title string, file *multipart.FileHeader) error {
	url, cover := utils.UploadVideo(file)
	//TODO 创建视频
	return videoDao.CreateVideo(&models.Video{Title: title, CoverUrl: cover, PlayUrl: url})
}

// FindVideoFeed 查询视频接口流
func FindVideoFeed(latest int64, limit int) (videoDTOList []common.VideoDTO, nextTime int64) {
	videoList := videoDao.GetVideoList(latest, limit)
	return createVideoDTOList(&videoList)
}

// PublishList 根据发布者查询视频
func PublishList(uid int64) (videoDTOList []common.VideoDTO) {
	videoList := videoDao.GetVideoListByUid(uid)
	videoDTOList, _ = createVideoDTOList(&videoList)
	return
}

// 构建视频dto切片
func createVideoDTOList(videoList *[]models.Video) (videoDTOList []common.VideoDTO, nextTime int64) {
	nextTime = (*videoList)[0].CreateTime
	for _, video := range *videoList {
		if video.CreateTime < nextTime {
			nextTime = video.CreateTime
		}
		videoDTOList = append(videoDTOList, createVideoDTO(video))
	}
	return
}

// 构建dto
func createVideoDTO(video models.Video) common.VideoDTO {
	userDTO, err := FindUserById(video.AuthorId)
	if err != nil {
		userDTO = &common.UserDTO{}
	}
	return common.VideoDTO{
		Id:            video.Id,
		Author:        *userDTO,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavouriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    videoDao.IsFavourite(common.GetUser(), video.Id),
		Title:         video.Title,
	}
}
