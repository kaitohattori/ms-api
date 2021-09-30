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

	APITimeoutSec time.Duration

	DbHost     string
	DbPort     int
	DbDriver   string
	DbName     string
	DbUser     string
	DbPassword string
	DbSslMode  string
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

// Config is variable of ConfigList
var Config ConfigList

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		LogFile:               cfg.Section("api").Key("log_file").String(),
		WebAPIHost:            cfg.Section("api").Key("web_api_host").String(),
		WebAPIPort:            cfg.Section("api").Key("web_api_port").MustInt(),
		StreamAPIHost:         cfg.Section("api").Key("stream_api_host").String(),
		StreamAPIPort:         cfg.Section("api").Key("stream_api_port").MustInt(),
		RecommendationAPIHost: cfg.Section("api").Key("recommendation_api_host").String(),
		RecommendationAPIPort: cfg.Section("api").Key("recommendation_api_port").MustInt(),
		APITimeoutSec:         time.Duration(cfg.Section("api").Key("timeout_sec").MustInt()) * time.Second,
		DbHost:                cfg.Section("db").Key("db_host").String(),
		DbPort:                cfg.Section("db").Key("db_port").MustInt(),
		DbDriver:              cfg.Section("db").Key("db_driver").String(),
		DbName:                cfg.Section("db").Key("db_name").String(),
		DbUser:                cfg.Section("db").Key("db_user").String(),
		DbPassword:            cfg.Section("db").Key("db_password").String(),
		DbSslMode:             cfg.Section("db").Key("db_ssl_mode").String(),
	}
}
