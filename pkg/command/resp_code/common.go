package resp_code

import "yang-backend/pkg/command/models"

var TimeParseFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "yang.00001", Message: "time parse filed"}}
var UploadImageFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "yang.00001", Message: "upload image filed"}}
