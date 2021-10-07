package main

import (
	"fmt"
	"ms-api/app/controller"
	"ms-api/app/repository"
	"ms-api/app/service"
	"ms-api/app/util"
	"ms-api/config"
	_ "ms-api/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	timeout "github.com/vearne/gin-timeout"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	StartServer()
}

func StartServer() {
	r := gin.Default()

	authUtil := util.NewAuthUtil(
		config.Config.Auth0Identifier,
		config.Config.Auth0Domain,
		config.Config.AuthHost,
	)
	r.Use(authUtil.CorsMiddleware())

	videoController := controller.NewVideoController()

	analysisRepository := repository.NewAnalysisRepository()
	analysisService := service.NewAnalysisService(analysisRepository)
	analysisController := controller.NewAnalysisController(analysisService)

	rateRepository := repository.NewRateRepository()
	rateService := service.NewRateService(rateRepository)
	rateController := controller.NewRateController(rateService)

	thumbnailController := controller.NewThumbnailController()

	v1 := r.Group("/api/v1")
	{
		videos := v1.Group("/videos")
		{
			videos.GET("", videoController.Find)
			videos.GET(":id", videoController.Get)
			videos.POST("", authUtil.CheckJWT(), videoController.Add)
			videos.POST(":id", authUtil.CheckJWT(), videoController.Update)
			videos.DELETE(":id", authUtil.CheckJWT(), videoController.Delete)
		}
		video_upload := v1.Group("/videos")
		{
			video_upload.POST("upload", authUtil.CheckJWT(), videoController.Upload)
		}
		analysis := v1.Group("/videos")
		{
			analysis.GET(":id/analysis/total", analysisController.Total)
			analysis.POST(":id/analysis", authUtil.CheckJWT(), analysisController.Add)
		}
		rates := v1.Group("/videos")
		{
			rates.GET(":id/rate", authUtil.CheckJWT(), rateController.Get)
			rates.GET(":id/rate/average", rateController.Average)
			rates.POST(":id/rate", authUtil.CheckJWT(), rateController.Update)
		}
		thumbnail := v1.Group("/videos")
		{
			thumbnail.GET(":id/thumbnail", thumbnailController.GetThumbnailImage)
		}
		videos.Use(
			timeout.Timeout(
				timeout.WithTimeout(config.Config.APITimeout),
				timeout.WithErrorHttpCode(http.StatusRequestTimeout),
				timeout.WithCallBack(func(r *http.Request) {
					fmt.Println("Request Timeout : ", r.URL.String())
				})),
		)
		video_upload.Use(
			timeout.Timeout(
				timeout.WithTimeout(config.Config.FileUploadAPITimeout),
				timeout.WithErrorHttpCode(http.StatusRequestTimeout),
				timeout.WithCallBack(func(r *http.Request) {
					fmt.Println("Request Timeout : ", r.URL.String())
				})),
		)
		analysis.Use(
			timeout.Timeout(
				timeout.WithTimeout(config.Config.APITimeout),
				timeout.WithErrorHttpCode(http.StatusRequestTimeout),
				timeout.WithCallBack(func(r *http.Request) {
					fmt.Println("Request Timeout : ", r.URL.String())
				})),
		)
		rates.Use(
			timeout.Timeout(
				timeout.WithTimeout(config.Config.APITimeout),
				timeout.WithErrorHttpCode(http.StatusRequestTimeout),
				timeout.WithCallBack(func(r *http.Request) {
					fmt.Println("Request Timeout : ", r.URL.String())
				})),
		)
		thumbnail.Use(
			timeout.Timeout(
				timeout.WithTimeout(config.Config.APITimeout),
				timeout.WithErrorHttpCode(http.StatusRequestTimeout),
				timeout.WithCallBack(func(r *http.Request) {
					fmt.Println("Request Timeout : ", r.URL.String())
				})),
		)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(fmt.Sprintf(":%d", config.Config.WebAPIPort))
}
