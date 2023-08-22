package service

import (
	"github.com/DzoniDiplomski/Backend_API/converter"
	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/model"
	"github.com/DzoniDiplomski/Backend_API/repo"
)

type RequisitionService struct {
}

var requisitionRepo = repo.RequisitionRepo{}

func (requisitionService *RequisitionService) CreateRequisition(requisition model.RequisitionDTO) error {
	id, err := requisitionRepo.Create(converter.FromDTORequisition(requisition))
	if err != nil {
		return err
	}

	err = addItemsToDBAndBindWithRequisition(id, requisition.Products)
	if err != nil {
		requisitionRepo.Delete(id)
		return err
	}
	return nil
}

func addItemsToDBAndBindWithRequisition(id int64, products []model.Product) error {
	tx, err := db.DBConn.Begin()
	if err != nil {
		return err
	}

	for _, product := range products {
		_, err := tx.Exec(db.PSCreateRequisitionItem, product.Name, product.Quantity, id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
