package server

import (
    "github.com/gin-gonic/gin"
    "yang-backend/pkg/book_keeping/route_func"
    "yang-backend/pkg/ginlog"
)

func RegisterRoute(r *gin.Engine)  {
    v1 := r.Group("v1", ginlog.Ginzap(ginlog.ZapLogger, "ebs", true), ginlog.RecoveryWithZap(ginlog.ZapLogger, true))
    testApiRegister(v1)
    bookKeepingRegister(v1)

}

func testApiRegister(rg *gin.RouterGroup)  {
    rg.GET("/test", func(context *gin.Context) {
        context.JSON(200, "test !!")
    })
}

func bookKeepingRegister(rg *gin.RouterGroup)  {
    rg.POST("/income", route_func.IncomeCreate)
    rg.GET("/incomes", route_func.IncomeGets)
    rg.POST("/income/category", route_func.IncomeCategoryCreate)
}