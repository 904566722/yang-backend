package main

import (
    "log"
    "yang-backend/pkg/config"
    "yang-backend/pkg/db"
    "yang-backend/pkg/ginlog"
    "yang-backend/pkg/server"
)
import "github.com/gin-gonic/gin"

func main() {
    config.InitConfig()
    r := gin.New()
    r.GET("/", func(context *gin.Context) {
        context.JSON(200, "hello")
    })

    ginlog.InitLogger(ginlog.Config{
        Env: "product",
        LogPath: config.Config.LogPath,
        MaxSize: 1024,
        MaxBackups: 5,
        MaxAge: 365,
        Compress: false,
    })
    db.InitDB()
    db.InitTable()
    server.RegisterRoute(r)


    err := r.Run(":8080")
    if err != nil {
        log.Fatalf("r.Run 8080 failed\n")
    }
}
