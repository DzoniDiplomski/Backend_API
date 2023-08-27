package repo

import (
	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/model"
)

type CalculationRepo struct {
}

func (calculationRepo *CalculationRepo) Create(calculaton model.Calculation) (int64, error) {
	var id int64
	err := db.DBConn.QueryRow(db.PSCreateCalculation).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
