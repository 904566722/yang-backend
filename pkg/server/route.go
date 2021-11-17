package server

import (
    "github.com/gin-gonic/gin"
    "yang-backend/pkg/ginlog"
)

func RegisterRoute(r *gin.Engine)  {
    v1 := r.Group("v1", ginlog.Ginzap(ginlog.ZapLogger, "ebs", true), ginlog.RecoveryWithZap(ginlog.ZapLogger, true))
    testApiRegister(v1)
}

func testApiRegister(rg *gin.RouterGroup)  {
    rg.GET("/test", func(context *gin.Context) {
        context.JSON(200, "test !!")
    })
}