package model

type RequisitionDTO struct {
	ManagerId int64     `json:"managerId"`
	Products  []Product `json:"products"`
}

type AllRequisitionsDTO struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type Requisition struct {
	ManagerId int64 `json:"managerId"`
}
