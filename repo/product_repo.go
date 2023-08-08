package repo

import (
	"diplomski.com/db"
	"diplomski.com/model"
)

type ProductRepo struct {
}

func (productRepo *ProductRepo) AddProducts(products []model.Product) error {
	tx, err := db.DBConn.Begin()
	if err != nil {
		return err
	}

	for _, p := range products {
		_, err := tx.Exec(db.PSAddProducts, p.Id, p.Barcode, p.Name)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
