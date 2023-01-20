package service

import (
	"TikTokApp/utils"
	"mime/multipart"
)

func Publish(title string, file *multipart.FileHeader) {
	url, cover := utils.UploadVideo(file)
	//	TODO 创建视频
}
