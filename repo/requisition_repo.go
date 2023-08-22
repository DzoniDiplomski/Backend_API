package repo

import (
	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/model"
)

type RequisitionRepo struct {
}

func (requisitionRepo *RequisitionRepo) Create(requisition model.Requisition) (int64, error) {
	var id int64
	err := db.DBConn.QueryRow(db.PSAddRequisition, requisition.ManagerId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (requisitionRepo *RequisitionRepo) Delete(id int64) error {
	_, err := db.DBConn.Exec(db.PSDeleteRequisition, id)
	if err != nil {
		return err
	}
	return nil
}
