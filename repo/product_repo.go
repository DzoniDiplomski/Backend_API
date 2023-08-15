package repo

import (
	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/model"
)

type ProductRepo struct {
}

func (productRepo *ProductRepo) AddProducts(products []model.Product) error {
	tx, err := db.DBConn.Begin()
	if err != nil {
		return err
	}

	for _, p := range products {
		_, err := tx.Exec(db.PSAddProducts, p.Id, p.Barcode, p.Name, 0)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
