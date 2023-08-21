package repo

import (
	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/model"
)

type ReceiptRepo struct {
}

func (receiptRepo *ReceiptRepo) Create(receipt model.Receipt, query string) (int64, error) {
	var receiptId int64
	err := db.DBConn.QueryRow(query, receipt.ShopId, receipt.CashBoxId).Scan(&receiptId)
	if err != nil {
		return 0, err
	}
	return receiptId, nil
}

func (receiptRepo *ReceiptRepo) Delete(id int64, query string) error {
	_, err := db.DBConn.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
