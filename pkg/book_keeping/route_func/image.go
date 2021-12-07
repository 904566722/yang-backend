package route_func

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"yang-backend/pkg/command/models"
	"yang-backend/pkg/command/resp_code"
	"yang-backend/pkg/config"
	"yang-backend/pkg/ginlog"
)

// 图片服务

type UploadImageOutput struct {
	models.ResponseBase
	Data ImageInfo `json:"data"`
}

type ImageInfo struct {
	Url string `json:"url"`
}

// UploadImage 单图片上传
func UploadImage(ctx *gin.Context) {
	image, err := ctx.FormFile("file")
	if err != nil {
		ginlog.CtxLogger(ctx).Error("get image content failed",
			zap.Error(err))
		ctx.JSON(200, resp_code.UploadImageFailed)
		return
	}
	//image.Filename
	//image.Size
	if err := ctx.SaveUploadedFile(image, config.Config.ImageUploadPath+image.Filename); err != nil {
		ctx.JSON(200, resp_code.UploadImageFailed)
		ginlog.CtxLogger(ctx).Error("save upload file failed",
			zap.Error(err))
		return
	}
	imageInfo := ImageInfo{
	    Url: config.Config.ImageUploadPath + image.Filename,
    }
    output := UploadImageOutput{
	    ResponseBase: models.Success,
	    Data: imageInfo,
    }
	ctx.JSON(200, output)
}
