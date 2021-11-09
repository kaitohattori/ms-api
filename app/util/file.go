package util

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"ms-api/config"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var FileUtil FileUtilFuncs

type FileUtilFuncs struct {
}

func init() {
	FileUtil = FileUtilFuncs{}
}

func FileUtilMakeWorkingDirectory(videoId int) (*string, error) {
	path := fmt.Sprintf("%s/%d", config.Config.AssetsWorkingDirPath, videoId)
	if err := os.MkdirAll(path, 0777); err != nil {
		return nil, err
	}
	return &path, nil
}

func FileUtilDeleteWorkingDirectory(dirPath string) error {
	return os.RemoveAll(dirPath)
}

func FileUtilSaveVideo(file multipart.File, header multipart.FileHeader, toDirPath string) (*string, error) {
	videoFilePath := fmt.Sprintf("%s/%s", toDirPath, config.Config.AssetsVideoFileName)
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

func FileUtilProcessVideo(ctx *gin.Context, videoId int, srcFilePath string) error {
	dstDirPath, _ := filepath.Abs(fmt.Sprintf("%s/%d", config.Config.AssetsDirPath, videoId))
	cmd := exec.CommandContext(ctx, "/bin/sh", config.Config.AssetsVideoProcessorScriptFilePath, srcFilePath, dstDirPath, "1920x1080")
	log.Println(cmd)
	// cmd.Start()
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func FileUtilThumbnailFilePath(videoId int) string {
	basePath := fmt.Sprintf("%s/%d", config.Config.AssetsDirPath, videoId)
	return fmt.Sprintf("%s/thumbnail/%s", basePath, config.Config.AssetsThumbnailFileName)
}
