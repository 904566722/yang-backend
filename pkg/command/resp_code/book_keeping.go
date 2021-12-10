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
var OutcomeDeleteFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0104", Message: "delete outcome failed"}}
var IncomeGetsFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0105", Message: "get incomes failed"}}
var OutcomeGetsFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0105", Message: "get outcomes failed"}}
var IncomeCountFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0106", Message: "count incomes failed"}}
var StatisticMonEatFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0106", Message: "statistic mon eat failed"}}


// income category 收入类别

var IncomeCategoryCreateFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0201", Message: "create income category failed"}}
var IncomeCategoryGetFailed  = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0202", Message: "get income category failed"}}
var OutcomeCategoryGetFailed  = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0202", Message: "get outcome category failed"}}
var IncomeCategoryUpdateFailed  = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0203", Message: "update income category failed"}}
var IncomeCategoryDeleteFailed  = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0204", Message: "delete income category failed"}}

var CtgGetsFailed  = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "book_keeping.0301", Message: "get categories failed"}}


var CreateFlowFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "create flow failed"}}
var CreateWaterFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "create water failed"}}
var CreateWaterCltFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "create water collection failed"}}
var CreateSortKlgFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "create sort knowledge failed"}}
var CreateTodoFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "create todo failed"}}
var CreateDevTestFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "create dev test failed"}}
var GetFlowsFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "get flows failed"}}
var GetWatersFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "get waters failed"}}
var UpdateWaterFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "update water failed"}}
var UpdateTodoFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "update todo failed"}}
var GetWaterFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "get water failed"}}
var GetSortKlgsFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "get sort klgs failed"}}
var GetTodosFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "get todos failed"}}
var GetWaterCltsFailed = models.ErrorResponse{ResponseBase: models.ResponseBase{Code: "flow.0101", Message: "get water clts failed"}}


