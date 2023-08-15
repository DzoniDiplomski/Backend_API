package model

type ReceiptDTO struct {
	CashBoxId int       `json:"kasaId"`
	ShopId    int       `json:"trafikaId"`
	CashierId int64     `json:"jmbgKasira"`
	Products  []Product `json:"products"`
}

type Receipt struct {
	CashBoxId int `json:"kasaId"`
	ShopId    int `json:"trafikaId"`
}
