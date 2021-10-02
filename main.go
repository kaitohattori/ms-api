package main

import (
	"fmt"
	"ms-api/app/controller"
	"ms-api/app/repository"
	"ms-api/app/service"
	"ms-api/config"
	_ "ms-api/docs"
	"net/http"
	"time"

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

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	r := gin.Default()

	videoRepository := repository.NewVideoRepository()
	videoService := service.NewVideoService(videoRepository)
	videoController := controller.NewVideoController(videoService)

	viewRepository := repository.NewViewRepository()
	viewService := service.NewViewService(viewRepository)
	viewController := controller.NewViewController(viewService)

	rateRepository := repository.NewRateRepository()
	rateService := service.NewRateService(rateRepository)
	rateController := controller.NewRateController(rateService)

	mediaRepository := repository.NewMediaRepository()
	mediaService := service.NewMediaService(mediaRepository)
	mediaController := controller.NewMediaController(mediaService)

	v1 := r.Group("/api/v1")
	{
		videos := v1.Group("/videos")
		{
			videos.GET("", videoController.Find)
			videos.GET(":id", videoController.Get)
			videos.POST("", videoController.Add)
			videos.POST(":id", videoController.Update)
			videos.DELETE(":id", videoController.Delete)
		}
		views := v1.Group("/videos")
		{
			views.GET(":id/view/total", viewController.Total)
			views.POST(":id/view", viewController.Add)
		}
		rates := v1.Group("/videos")
		{
			rates.GET(":id/rate", rateController.Get)
			rates.GET(":id/rate/average", rateController.Average)
			rates.POST(":id/rate", rateController.Add)
		}
		media := v1.Group("/videos")
		{
			media.GET(":id/thumbnail", mediaController.GetThumbnailImage)
			media.POST("upload", mediaController.Upload)
		}
		videos.Use(
			timeout.Timeout(
				timeout.WithTimeout(config.Config.APITimeoutSec),
				timeout.WithErrorHttpCode(http.StatusRequestTimeout),
				timeout.WithCallBack(func(r *http.Request) {
					fmt.Println("Request Timeout : ", r.URL.String())
				})),
		)
		views.Use(
			timeout.Timeout(
				timeout.WithTimeout(config.Config.APITimeoutSec),
				timeout.WithErrorHttpCode(http.StatusRequestTimeout),
				timeout.WithCallBack(func(r *http.Request) {
					fmt.Println("Request Timeout : ", r.URL.String())
				})),
		)
		rates.Use(
			timeout.Timeout(
				timeout.WithTimeout(config.Config.APITimeoutSec),
				timeout.WithErrorHttpCode(http.StatusRequestTimeout),
				timeout.WithCallBack(func(r *http.Request) {
					fmt.Println("Request Timeout : ", r.URL.String())
				})),
		)
		media.Use(
			timeout.Timeout(
				timeout.WithTimeout(60*time.Second),
				timeout.WithErrorHttpCode(http.StatusRequestTimeout),
				timeout.WithCallBack(func(r *http.Request) {
					fmt.Println("Request Timeout : ", r.URL.String())
				})),
		)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
