package route_func

import (
    "github.com/gin-gonic/gin"
    "time"
    "yang-backend/pkg/command/models"
    "yang-backend/pkg/command/resp_code"
    "yang-backend/pkg/command/utils"
    "yang-backend/pkg/config"
)

type GetCurYearMonOutput struct {
	models.ResponseBase
	Data YearMon `json:"data"`
}

type YearMon struct {
    Year int `json:"year"`
    Mon int `json:"mon"`
}

func GetCurYearMon(ctx *gin.Context) {
    var y, m int
	curTime := ctx.Query("cur_time")
	if curTime != "" {
        now, err := utils.ParseTime(ctx, config.TimeLayoutStr, curTime)
        if err != nil {
            ctx.JSON(200, resp_code.TimeParseFailed)
            return
        }
        y, m = config.GetMon(now)
    } else {
        now := time.Now()
        y, m = config.GetMon(now)
    }
    ctx.JSON(200, GetCurYearMonOutput{
        ResponseBase: models.Success,
        Data: YearMon{
            Year: y,
            Mon: m,
        },
    })
}
