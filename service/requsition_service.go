package service

import (
	"time"

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

func (requisitionService *RequisitionService) GetRequisitions(offset int, limit int) ([]model.AllRequisitionsDTO, error) {
	rows, err := db.DBConn.Query(db.PSGetAllRequisitions, limit, offset)
	if err != nil {
		return nil, err
	}

	var requisitions []model.AllRequisitionsDTO
	for rows.Next() {
		var requisition model.AllRequisitionsDTO
		err := rows.Scan(&requisition.Id, &requisition.CreatedAt)
		if err != nil {
			return nil, err
		}
		createdAtTime, err := time.Parse(time.RFC3339, requisition.CreatedAt)
		if err != nil {
			return nil, err
		}

		formattedTime := createdAtTime.Format("02-01-2006 15:04")
		requisition.CreatedAt = formattedTime
		requisitions = append(requisitions, requisition)
	}
	return requisitions, nil
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

func (requisitionService *RequisitionService) GetRequisitionItems(id int64) ([]model.ProductRequisitionDTO, error) {
	rows, err := db.DBConn.Query(db.PSGetRequisitionItems, id)
	if err != nil {
		return nil, err
	}

	var items []model.ProductRequisitionDTO
	for rows.Next() {
		var item model.ProductRequisitionDTO
		err = rows.Scan(&item.Quantity, &item.Name)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (requisitionService *RequisitionService) CalculatePagesForAllRequisitions(itemsPerPage int) (model.AllReceiptsPages, error) {
	var count int

	err := db.DBConn.QueryRow(db.PSCountAllRequisitions).Scan(&count)
	if err != nil {
		return model.AllReceiptsPages{}, err
	}

	numberOfPages := count / itemsPerPage
	if numberOfPages != 0 {
		leftoverItems := count % itemsPerPage
		return model.AllReceiptsPages{
			NumberOfPages: numberOfPages,
			LeftoverItems: leftoverItems,
		}, nil
	}

	numberOfPages++
	return model.AllReceiptsPages{
		NumberOfPages: numberOfPages,
		LeftoverItems: 0,
	}, nil
}
