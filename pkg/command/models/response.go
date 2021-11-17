package models

type ResponseBase struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

var Success = ResponseBase{Code: "yang.0000", Message: "Success"}

type ErrorResponse struct {
    ResponseBase
    Err error `json:"-"`
}

func (resp *ResponseBase) HttpCode() int {
    return 200
}

func InputError(err error) *ErrorResponse {
    return &ErrorResponse{ResponseBase: ResponseBase{Code: "yang.0101", Message: "input error"}, Err: err}
}

func InternalError(err error) *ErrorResponse {
    return &ErrorResponse{ResponseBase: ResponseBase{Code: "yang.0001", Message: "Internal error "}, Err: err}
}

// LogErrAndResp 打印错误日志并返回错误信息
//func LogErrAndResp(ctx *gin.Context, err error)  {
//    errResp, ok := err.(*ErrorResponse)
//    if !ok {
//        errResp = InternalError(err)
//    }
//    if errResp.Err != nil {
//
//    }
//}