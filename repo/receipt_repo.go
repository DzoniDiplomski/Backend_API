package repo

import (
	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/model"
)

type ReceiptRepo struct {
}

func (receiptRepo *ReceiptRepo) Create(receipt model.Receipt) (int64, error) {
	var receiptId int64
	err := db.DBConn.QueryRow(db.PSAddReceipt, receipt.ShopId, receipt.CashBoxId).Scan(&receiptId)
	if err != nil {
		return 0, err
	}
	return receiptId, nil
}

func (receiptRepo *ReceiptRepo) Delete(id int64) error {
	_, err := db.DBConn.Exec(db.PSDeleteReceipt, id)
	if err != nil {
		return err
	}
	return nil
}
