package service

import (
	"database/sql"

	"github.com/DzoniDiplomski/Backend_API/converter"
	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/model"
	"github.com/DzoniDiplomski/Backend_API/repo"
)

type ReceiptService struct {
}

var receiptRepo = &repo.ReceiptRepo{}

func (receiptService *ReceiptService) CreateReceipt(receipt model.ReceiptDTO) error {
	tx, err := db.DBConn.Begin()
	if err != nil {
		return err
	}

	receiptId, err := receiptRepo.Create(converter.FromDTOReceipt(receipt))
	if err != nil {
		tx.Rollback()
		return err
	}

	err = addItemsToDBAndBindWithReceipt(receipt, tx, receiptId)
	if err != nil {
		receiptRepo.Delete(receiptId)
		tx.Rollback()
		return err
	}

	err = bindReceiptWithCashier(tx, receiptId, receipt)
	if err != nil {
		receiptRepo.Delete(receiptId)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func bindReceiptWithCashier(tx *sql.Tx, receiptId int64, receipt model.ReceiptDTO) error {
	_, err := tx.Exec(db.PSBindReceiptWithCashier, receiptId, receipt.CashierId)
	if err != nil {
		return err
	}
	return nil
}

func addItemsToDBAndBindWithReceipt(receipt model.ReceiptDTO, tx *sql.Tx, receiptId int64) error {
	for _, product := range receipt.Products {
		var itemId int64
		err := tx.QueryRow(db.PSCreateReceiptItem, product.Quantity, product.Id, product.Price).Scan(&itemId)
		if err != nil {
			return err
		}

		_, err = tx.Exec(db.PSBindItemWithReceipt, receiptId, itemId)
		if err != nil {
			return err
		}
	}
	return nil
}
