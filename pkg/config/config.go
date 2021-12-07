package config

import (
	"io"
	"log"
	"os"
	"time"
)

type config struct {
	LogPath         string
	DbUrl           string
	DbUser          string
	DbPw            string
	DbPort          string
	DbName          string
	DbPath          string
	EnableDbLog     bool
	ImageUploadPath string
}

const (
	EnvKeyLogPath         = "LOG_PATH"
	EnvKeyDBURL           = "DB_URL"
	EnvKeyDBUSER          = "DB_USER"
	EnvKeyDBPW            = "DB_PW"
	EnvKeyDBPORT          = "DB_PORT"
	EnvKeyDBNAME          = "DB_NAME"
	EnvKeyEnableDbLog     = "ENABLE_DB_LOG"
	EnvKeyImageUploadPath = "IMAGE_UPLOAD_PATH"
)

const (
	TimeLayoutStr = "2006-01-02 15:04:05"
)

var NilTime = time.Time{}

var Config config

func InitConfig() {
	if Config.LogPath = os.Getenv(EnvKeyLogPath); Config.LogPath == "" {
		Config.LogPath = "/var/log/yang/yang.log"
	}
	logFile, err := os.OpenFile(Config.LogPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("os.OpenFile failed. logPath: %s\n", Config.LogPath)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	if Config.DbUrl = os.Getenv(EnvKeyDBURL); Config.DbUrl == "" {
		log.Fatalf("get env %s failed\n", EnvKeyDBURL)
	}
	if Config.DbUser = os.Getenv(EnvKeyDBUSER); Config.DbUser == "" {
		log.Fatalf("get env %s failed\n", EnvKeyDBUSER)
	}
	if Config.DbPw = os.Getenv(EnvKeyDBPW); Config.DbPw == "" {
		log.Fatalf("get env %s failed\n", EnvKeyDBPW)
	}
	if Config.DbPort = os.Getenv(EnvKeyDBPORT); Config.DbPort == "" {
		log.Fatalf("get env %s failed\n", EnvKeyDBPORT)
	}
	if Config.DbName = os.Getenv(EnvKeyDBNAME); Config.DbUser == "" {
		log.Fatalf("get env %s failed\n", EnvKeyDBNAME)
	}
	Config.DbPath = Config.DbUser + ":" + Config.DbPw + "@tcp(" + Config.DbUrl + ":" + Config.DbPort + ")/" + Config.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Printf("database path: %s\n", Config.DbPath)

	if enableDbLog := os.Getenv(EnvKeyEnableDbLog); enableDbLog == "" {
		log.Fatalf("get env %s failed\n", EnvKeyEnableDbLog)
	} else if enableDbLog == "true" {
		Config.EnableDbLog = true
	} else if enableDbLog == "false" {
		Config.EnableDbLog = false
	}

	if Config.ImageUploadPath = os.Getenv(EnvKeyImageUploadPath); Config.ImageUploadPath == "" {
		log.Fatalf("get env %s failed\n", EnvKeyImageUploadPath)
	}
}
