package model

type Product struct {
	Id       int     `json:"sif"`
	Barcode  string  `json:"bc"`
	Name     string  `json:"naz"`
	Price    float64 `json:"cena"`
	Quantity int     `json:"kolicina"`
}
