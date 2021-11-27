package main

import (
    "log"
    "net/http"
    "yang-backend/pkg/config"
    "yang-backend/pkg/db"
    "yang-backend/pkg/ginlog"
    "yang-backend/pkg/server"
)
import "github.com/gin-gonic/gin"

func main() {
    config.InitConfig()
    r := gin.New()
    r.Use(Cors())   //  解决跨域问题
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


    err := r.Run(":8888")
    if err != nil {
        log.Fatalf("r.Run 8080 failed\n")
    }
}

func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method
        origin := c.Request.Header.Get("Origin")
        if origin != "" {
            c.Header("Access-Control-Allow-Origin", "*")
            c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
            c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Access-Token, x-token")
            c.Header("Access-Control-Allow-Credentials", "true")
            c.Set("content-type", "application/json")
        }
        //放行所有OPTIONS方法
        if method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
        }
        c.Next()
    }
}