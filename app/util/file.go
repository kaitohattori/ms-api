package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var FileUtil FileUtilFuncs

type FileUtilFuncs struct {
	MediaRoot                    string
	MediaWorkingRoot             string
	VideoFileName                string
	ThumbnailFileName            string
	VideoProcessorScriptFilePath string
}

func init() {
	FileUtil = FileUtilFuncs{
		MediaRoot:                    "assets/media",
		MediaWorkingRoot:             "assets/working",
		VideoFileName:                "video.mp4",
		ThumbnailFileName:            "thumbnail.jpg",
		VideoProcessorScriptFilePath: "scripts/upload_video.sh",
	}
}

func (u FileUtilFuncs) MakeWorkingDirectory(videoId int) (*string, error) {
	path := fmt.Sprintf("%s/%d", u.MediaWorkingRoot, videoId)
	if err := os.MkdirAll(path, 0777); err != nil {
		return nil, err
	}
	return &path, nil
}

func (u FileUtilFuncs) DeleteWorkingDirectory(dirPath string) error {
	return os.RemoveAll(dirPath)
}

func (u FileUtilFuncs) SaveVideo(file multipart.File, header multipart.FileHeader, toDirPath string) (*string, error) {
	videoFilePath := fmt.Sprintf("%s/%s", toDirPath, u.VideoFileName)
	newFile, err := os.Create(videoFilePath)
	if err != nil {
		return nil, err
	}
	defer newFile.Close()
	if _, err := io.Copy(newFile, file); err != nil {
		return nil, err
	}
	absVideoFilePath, err := filepath.Abs(videoFilePath)
	if err != nil {
		return nil, err
	}
	return &absVideoFilePath, nil
}

func (u FileUtilFuncs) ProcessVideo(ctx *gin.Context, videoId int, srcFilePath string) error {
	dstDirPath, _ := filepath.Abs(fmt.Sprintf("%s/%d", u.MediaRoot, videoId))
	cmd := exec.CommandContext(ctx, "/bin/sh", u.VideoProcessorScriptFilePath, srcFilePath, dstDirPath, "1920x1080")
	fmt.Println(cmd)
	// cmd.Start()
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (u FileUtilFuncs) ThumbnailFilePath(videoId int) string {
	basePath := fmt.Sprintf("%s/%d", u.MediaRoot, videoId)
	return fmt.Sprintf("%s/thumbnail/%s", basePath, u.ThumbnailFileName)
}
