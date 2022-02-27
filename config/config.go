package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	LogFile string

	WebAPIHost            string
	WebAPIPort            int
	StreamAPIHost         string
	StreamAPIPort         int
	RecommendationAPIHost string
	RecommendationAPIPort int

	APITimeout           time.Duration
	FileUploadAPITimeout time.Duration

	AssetsDirPath                      string
	AssetsWorkingDirPath               string
	AssetsVideoFileName                string
	AssetsThumbnailFileName            string
	AssetsVideoProcessorScriptFilePath string

	DbHost     string
	DbPort     int
	DbDriver   string
	DbName     string
	DbUser     string
	DbPassword string
	DbSslMode  string

	Auth0Domain       string
	Auth0Identifier   string
	Auth0ClientId     string
	Auth0ClientSecret string
}

func (c ConfigList) WebAPIURL() string {
	return fmt.Sprintf("http://%s:%d", c.WebAPIHost, c.WebAPIPort)
}

func (c ConfigList) StreamAPIURL() string {
	return fmt.Sprintf("http://%s:%d", c.StreamAPIHost, c.StreamAPIPort)
}

func (c ConfigList) RecommendationAPIURL() string {
	return fmt.Sprintf("http://%s:%d", c.RecommendationAPIHost, c.RecommendationAPIPort)
}

func getEnv() string {
	return os.Getenv("APP_ENV")
}

// Config is variable of ConfigList
var Config ConfigList

func init() {
	configFilePath := "config/config-production.ini"
	if getEnv() == "development" {
		configFilePath = "config/config-development.ini"
	}
	cfg, err := ini.Load(configFilePath)
	if err != nil {
		log.Fatalln("Failed to read file: ", err)
		os.Exit(1)
	}

	Config = ConfigList{
		LogFile:                            cfg.Section("api").Key("log_file").String(),
		WebAPIHost:                         cfg.Section("api").Key("web_api_host").String(),
		WebAPIPort:                         cfg.Section("api").Key("web_api_port").MustInt(),
		StreamAPIHost:                      cfg.Section("api").Key("stream_api_host").String(),
		StreamAPIPort:                      cfg.Section("api").Key("stream_api_port").MustInt(),
		RecommendationAPIHost:              cfg.Section("api").Key("recommendation_api_host").String(),
		RecommendationAPIPort:              cfg.Section("api").Key("recommendation_api_port").MustInt(),
		APITimeout:                         time.Duration(cfg.Section("api").Key("api_timeout_sec").MustInt()) * time.Second,
		AssetsDirPath:                      cfg.Section("file").Key("dir_path").String(),
		AssetsWorkingDirPath:               cfg.Section("file").Key("working_dir_path").String(),
		AssetsVideoFileName:                cfg.Section("file").Key("video_filename").String(),
		AssetsThumbnailFileName:            cfg.Section("file").Key("thumbnail_filename").String(),
		AssetsVideoProcessorScriptFilePath: cfg.Section("file").Key("video_processor_script_file_path").String(),
		DbHost:                             cfg.Section("db").Key("db_host").String(),
		DbPort:                             cfg.Section("db").Key("db_port").MustInt(),
		DbDriver:                           cfg.Section("db").Key("db_driver").String(),
		DbName:                             cfg.Section("db").Key("db_name").String(),
		DbUser:                             cfg.Section("db").Key("db_user").String(),
		DbPassword:                         cfg.Section("db").Key("db_password").String(),
		DbSslMode:                          cfg.Section("db").Key("db_ssl_mode").String(),
		Auth0Domain:                        cfg.Section("auth0").Key("domain").String(),
		Auth0Identifier:                    cfg.Section("auth0").Key("identifier").String(),
		Auth0ClientId:                      cfg.Section("auth0").Key("client_id").String(),
		Auth0ClientSecret:                  cfg.Section("auth0").Key("client_secret").String(),
	}
}
