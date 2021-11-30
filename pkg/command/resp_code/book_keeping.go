package resp_code

import "yang-backend/pkg/command/models"

var NotFoundResource = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "yang.00001", Message: "not found resource"}}

// income 收入

var IncomeCreateFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0101", Message: "create income failed"}}
var OutcomeCreateFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0101", Message: "create outcome failed"}}
var IncomeGetFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0102", Message: "get income failed"}}
var OutcomeGetFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0102", Message: "get outcome failed"}}
var IncomeUpdateFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0103", Message: "update income failed"}}
var OutcomeUpdateFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0103", Message: "update outcome failed"}}
var IncomeDeleteFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0104", Message: "delete income failed"}}
var IncomeGetsFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0105", Message: "get incomes failed"}}
var IncomeCountFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0106", Message: "count incomes failed"}}


// income category 收入类别

var IncomeCategoryCreateFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0201", Message: "create income category failed"}}
var IncomeCategoryGetFailed  = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0202", Message: "get income category failed"}}
var OutcomeCategoryGetFailed  = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0202", Message: "get outcome category failed"}}
var IncomeCategoryUpdateFailed  = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0203", Message: "update income category failed"}}
var IncomeCategoryDeleteFailed  = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0204", Message: "delete income category failed"}}

var CtgGetsFailed  = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0301", Message: "get categories failed"}}
