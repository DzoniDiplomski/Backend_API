package service

import (
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
		return err
	}

}
