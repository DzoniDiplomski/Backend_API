package model

type Calculation struct {
	Id       int64                   `json:"id"`
	Provider string                  `json:"provider"`
	Items    []CalculationProductDTO `json:"items"`
}
