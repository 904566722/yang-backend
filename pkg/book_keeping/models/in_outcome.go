package models

type CtgList struct {
	IncomeCtgs []IncomeCategory `json:"income_ctgs"`
	OutcomeCtgs []OutcomeCategory `json:"outcome_ctgs"`
}
