package util

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

var MediaUtil = MediaUtilFuncs{}

type MediaUtilFuncs struct{}

func (_ MediaUtilFuncs) MakeWorkingDirectory(videoId int) (*string, error) {
	path := fmt.Sprintf("assets/tmp/%d", videoId)
	if err := os.MkdirAll(path, 0777); err != nil {
		return nil, err
	}
	return &path, nil
}

func (_ MediaUtilFuncs) DeleteWorkingDirectory(dirPath string) error {
	return os.RemoveAll(dirPath)
}

func (_ MediaUtilFuncs) SaveVideo(fileName string, toDirPath string) error {
	newFile, _ := os.Create(fmt.Sprintf("%s/video.mp4", toDirPath))
	defer newFile.Close()
	file, _ := os.Open(fileName)
	defer file.Close()
	if _, err := io.Copy(newFile, file); err != nil {
		return err
	}
	return nil
}

func (_ MediaUtilFuncs) VideoProcessorFilePath() string {
	shellFilePath, _ := filepath.Abs("script/process_video.sh")
	return shellFilePath
}

func (u MediaUtilFuncs) ProcessVideo(ctx *gin.Context, videoId int, dirPath string) error {
	shellFilePath := u.VideoProcessorFilePath()
	cmd := exec.CommandContext(ctx, "/bin/sh", shellFilePath, dirPath, "1920x1080", strconv.Itoa(int(videoId)))
	// cmd.Start()
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
