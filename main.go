package main

import (
	"log"
	"ms-api/app/controller"
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
	util.LoggingSettings(config.Config.LogFile)
	StartServer()
}

func StartServer() {
	engine := gin.Default()

	// Auth
	authUtil := util.NewAuthUtil(
		config.Config.Auth0Identifier,
		config.Config.Auth0Domain,
	)

	// Timeout
	engine.Use(
		timeout.Timeout(
			timeout.WithTimeout(config.Config.APITimeout),
			timeout.WithErrorHttpCode(http.StatusRequestTimeout),
			timeout.WithCallBack(func(r *http.Request) {
				log.Println("Request Timeout : ", r.URL.String())
			})),
	)

	// Controller
	authController := controller.NewAuthController(authUtil)
	videoController := controller.NewVideoController()
	analysisController := controller.NewAnalysisController()
	rateController := controller.NewRateController()
	thumbnailController := controller.NewThumbnailController()

	// Router
	v1 := engine.Group("/api/v1")
	{
		auth := v1.Group("/")
		{
			auth.GET("login", authController.Login)
		}
		videoUpload := v1.Group("/videos")
		{
			videoUpload.POST("upload", authUtil.CheckJWT(), videoController.Upload)
		}
		videos := v1.Group("/videos")
		{
			videos.GET("", authUtil.CheckJWTAuthNotRequired(), videoController.Find)
			videos.GET(":id", videoController.Get)
			videos.POST("", authUtil.CheckJWT(), videoController.Add)
			videos.POST(":id", authUtil.CheckJWT(), videoController.Update)
			videos.DELETE(":id", authUtil.CheckJWT(), videoController.Delete)
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
			rates.PATCH(":id/rate", authUtil.CheckJWT(), rateController.Update)
		}
		thumbnail := v1.Group("/videos")
		{
			thumbnail.GET(":id/thumbnail", thumbnailController.GetThumbnail)
		}
	}
	engine.GET("/docs/api/v1/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.Run(":8080")
}
