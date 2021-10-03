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

var MediaUtil = MediaUtilFuncs{}

type MediaUtilFuncs struct{}

func (u MediaUtilFuncs) MakeWorkingDirectory(videoId int) (*string, error) {
	path := fmt.Sprintf("assets/tmp/%d", videoId)
	if err := os.MkdirAll(path, 0777); err != nil {
		return nil, err
	}
	return &path, nil
}

func (u MediaUtilFuncs) DeleteWorkingDirectory(dirPath string) error {
	return os.RemoveAll(dirPath)
}

func (u MediaUtilFuncs) SaveVideo(file multipart.File, header multipart.FileHeader, toDirPath string) (*string, error) {
	videoFilePath := fmt.Sprintf("%s/video.mp4", toDirPath)
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

func (u MediaUtilFuncs) VideoProcessorFilePath() string {
	shellFilePath, _ := filepath.Abs("scripts/process_video.sh")
	return shellFilePath
}

func (u MediaUtilFuncs) ProcessVideo(ctx *gin.Context, videoId int, srcFilePath string) error {
	shellFilePath := u.VideoProcessorFilePath()
	dstDirPath, _ := filepath.Abs(fmt.Sprintf("assets/media/%d", videoId))
	cmd := exec.CommandContext(ctx, "/bin/sh", shellFilePath, srcFilePath, dstDirPath, "1920x1080")
	fmt.Println(cmd)
	// cmd.Start()
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
