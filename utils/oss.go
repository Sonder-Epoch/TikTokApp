package utils

import (
	"TikTokApp/config"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"mime/multipart"
	"time"
)

var OSS *oss.Bucket

func InitOSS() {
	ossConfig := config.Conf.OSS
	client, err := oss.New(ossConfig.Endpoint, ossConfig.AccessKeyId, ossConfig.AccessKeySecret)
	if err != nil {
		panic("创建oss链接失败" + err.Error())
	}
	bucket, err := client.Bucket(ossConfig.BucketName)
	if err != nil {
		panic("不存在的存储空间" + err.Error())
	}
	OSS = bucket
}

// UploadVideo 上传视频,返回路径与封面
func UploadVideo(file *multipart.FileHeader) (url string, cover string) {
	open, err := file.Open()
	if err != nil {
		zap.L().Sugar().Errorf("视频传输失败%v", err)
		return
	}
	filename := file.Filename
	videoName := fmt.Sprintf("video/%v/%v", time.Now().Format("2006/01/02"), filename)
	err = OSS.PutObject(videoName, open)
	if err != nil {
		zap.L().Sugar().Errorf("视频存储失败%v", err)
		return
	}
	url = fmt.Sprintf("%v%v.%v/%v", "https://", config.Conf.OSS.BucketName, config.Conf.OSS.Endpoint, videoName)
	cover = fmt.Sprintf("%v?x-oss-process=video/snapshot,t_2000,f_jpg,w_800,h_600,m_fast", url)
	return
}
