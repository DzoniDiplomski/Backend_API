package model

type Product struct {
	Id       int     `json:"sif"`
	Barcode  string  `json:"bc"`
	Name     string  `json:"naz"`
	Price    float64 `json:"cena"`
	Quantity int     `json:"kolicina"`
}

type ProductDTO struct {
	Id        int     `json:"sif"`
	Price     float64 `json:"cena"`
	StartDate string  `json:"pocetak_vazenja"`
	EndDate   string  `json:"kraj_vazenja"`
}
