package model

type RequisitionDTO struct {
	ManagerId int64     `json:"managerId"`
	Products  []Product `json:"products"`
}

type Requisition struct {
	ManagerId int64 `json:"managerId"`
}
