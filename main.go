package main

import (
	"fmt"
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

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.Config.AuthHost)
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func StartServer() {
	engine := gin.Default()

	// Auth
	authUtil := util.NewAuthUtil(
		config.Config.Auth0Identifier,
		config.Config.Auth0Domain,
	)
	engine.Use(CorsMiddleware())

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
	videoController := controller.NewVideoController()
	analysisController := controller.NewAnalysisController()
	rateController := controller.NewRateController()
	thumbnailController := controller.NewThumbnailController()

	// Router
	v1 := engine.Group("/api/v1")
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
			thumbnail.GET(":id/thumbnail", thumbnailController.GetThumbnail)
		}
	}
	engine.GET(fmt.Sprintf("%s/*any", config.Config.ApiDocsPath), ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.Run(fmt.Sprintf(":%d", config.Config.WebAPIPort))
}
