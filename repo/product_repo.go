package repo

import (
	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/model"
)

type ProductRepo struct {
}

func (productRepo *ProductRepo) AddProducts(products []model.Product) ([]int64, error) {
	tx, err := db.DBConn.Begin()
	if err != nil {
		return nil, err
	}

	var (
		productId  int64
		productIds []int64
	)
	for _, p := range products {

		err := tx.QueryRow(db.PSAddProducts, p.Id, p.Barcode, p.Name, 0).Scan(&productId)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		productIds = append(productIds, productId)
	}

	tx.Commit()
	return productIds, nil
}
