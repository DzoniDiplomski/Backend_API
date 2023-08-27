package service

import (
	"database/sql"

	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/model"
	"github.com/DzoniDiplomski/Backend_API/repo"
)

type CalculationService struct {
}

var calculationRepo = repo.CalculationRepo{}

func (calculationService *CalculationService) CreateCalculation(calculation model.Calculation) error {
	tx, err := db.DBConn.Begin()
	if err != nil {
		return err
	}

	calculationId, err := calculationRepo.Create(calculation)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = createCalculationItems(tx, calculation.Items, calculationId)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func createCalculationItems(tx *sql.Tx, items []model.CalculationProductDTO, calculationId int64) error {
	for _, item := range items {
		var itemId int64
		err := tx.QueryRow(db.PSCreateCalculationItem, item.Id, item.Neto, item.Margin, item.Quantity, item.Pdv).Scan(&itemId)
		if err != nil {
			return err
		}

		_, err = tx.Exec(db.PSBindCalculationItemWithCalculation, calculationId, itemId)
		if err != nil {
			return err
		}

		_, err = tx.Exec(db.PSIncreaseItemQuantity, item.Quantity, item.Id)
		if err != nil {
			return err
		}
	}
	return nil
}
