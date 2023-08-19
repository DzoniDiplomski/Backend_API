package model

type ReceiptDTO struct {
	CashBoxId int       `json:"kasaId"`
	ShopId    int       `json:"trafikaId"`
	CashierId int64     `json:"jmbgKasira"`
	Products  []Product `json:"products"`
}

type AllReceiptsDTO struct {
	Id              int64  `json:"id"`
	ShopName        string `json:"nazivTrafike"`
	CashierName     string `json:"nazivProdavca"`
	CashierLastName string `json:"prezimeProdavca"`
	CreatedAt       string `json:"createdAt"`
}

type Receipt struct {
	CashBoxId int `json:"kasaId"`
	ShopId    int `json:"trafikaId"`
}

type AllReceiptsPages struct {
	NumberOfPages int `json:"numberOfPages"`
	LeftoverItems int `json:"leftoverItems"`
}
