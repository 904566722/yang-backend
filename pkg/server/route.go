package server

import (
    "github.com/gin-gonic/gin"
    "yang-backend/pkg/book_keeping/route_func"
    "yang-backend/pkg/command/models"
    "yang-backend/pkg/ginlog"
)

func RegisterRoute(r *gin.Engine)  {
    v1 := r.Group("v1", ginlog.Ginzap(ginlog.ZapLogger, "ebs", true), ginlog.RecoveryWithZap(ginlog.ZapLogger, true))
    testApiRegister(v1)
    bookKeepingRegister(v1)
    commonRegister(v1)
}

type LoginOutput struct {
    models.ResponseBase
    Token string `json:"token"`
}

type UserInfoOutput struct {
    models.ResponseBase
    Data UserInfo `json:"data"`
}

type UserInfo struct {
    Roles []string `json:"roles"`
    Name string `json:"name"`
    Avatar string `json:"avatar"`
    Introduction string `json:"introduction"`
}

func testApiRegister(rg *gin.RouterGroup)  {
    rg.GET("/test", func(context *gin.Context) {
        context.JSON(200, "yang project!")
    })
    rg.GET("/user/login", func(ctx *gin.Context) {
        ctx.JSON(200, LoginOutput{
            ResponseBase: models.Success,
            Token: "124364765857dvxvx",
        })
    })
    rg.GET("/user/info", func(ctx *gin.Context) {
        ctx.JSON(200, UserInfoOutput{
          ResponseBase: models.Success,
          Data: UserInfo{
              Roles: []string{"admin"},
              Name: "admin",
              Avatar: "https://avatar-static.segmentfault.com/732/311/732311852-5d5287d4a718b_huge128",
              Introduction: "旸",
          },
        })
    })
}

// bookKeepingRegister 注册 记账模块 的api
func bookKeepingRegister(rg *gin.RouterGroup)  {
    rg.POST("/income", route_func.IncomeCreate)
    rg.POST("/outcome", route_func.CreateOutcome)
    rg.GET("/income/:income_id", route_func.IncomeGet)
    rg.POST("/income/:income_id", route_func.IncomeUpdate)
    rg.POST("/outcome/:outcome_id", route_func.UpdateOutcome)
    rg.DELETE("/income/:income_id", route_func.IncomeDelete)
    rg.DELETE("/outcome/:outcome_id", route_func.DeleteOutcome)
    rg.GET("/incomes", route_func.IncomeGets)

    rg.POST("/income/category", route_func.IncomeCategoryCreate)
    rg.GET("/income/category/:income_category_id", route_func.IncomeCategoryGet)
    rg.POST("/income/categories", route_func.IncomeCategoryGets)
    rg.POST("/outcome/categories", route_func.GetOutcomeCategories)
    rg.POST("/income/category/:income_category_id", route_func.IncomeCategoryUpdate)
    rg.DELETE("/income/category/:income_category_id", route_func.IncomeCategoryDelete)

    rg.POST("/category/:op_unit", route_func.GetCtgs)
}

func commonRegister(rg *gin.RouterGroup)  {
    rg.GET("/year-mon", route_func.GetCurYearMon)
}