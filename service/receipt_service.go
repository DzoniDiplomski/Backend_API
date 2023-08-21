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

	var (
		createQuery      string
		deleteQuery      string
		bindItemQuery    string
		bindCashierQuery string
	)

	if receipt.EIN != 0 {
		createQuery = db.PSAddInvoice
		bindItemQuery = db.PSBindItemWithInvoice
		bindCashierQuery = db.PSBindInvoiceWithCashier
		deleteQuery = db.PSDeleteInvoice
	} else {
		createQuery = db.PSAddReceipt
		bindItemQuery = db.PSBindItemWithReceipt
		bindCashierQuery = db.PSBindReceiptWithCashier
		deleteQuery = db.PSDeleteReceipt
	}

	receiptId, err := receiptRepo.Create(converter.FromDTOReceipt(receipt), createQuery)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = addItemsToDBAndBindWithReceipt(receipt, tx, receiptId, bindItemQuery)
	if err != nil {
		receiptRepo.Delete(receiptId, deleteQuery)
		tx.Rollback()
		return err
	}

	err = bindReceiptWithCashier(tx, receiptId, receipt, bindCashierQuery)
	if err != nil {
		receiptRepo.Delete(receiptId, deleteQuery)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func bindReceiptWithCashier(tx *sql.Tx, receiptId int64, receipt model.ReceiptDTO, query string) error {
	_, err := tx.Exec(query, receiptId, receipt.CashierId, receipt.EIN)
	if err != nil {
		return err
	}
	return nil
}

func addItemsToDBAndBindWithReceipt(receipt model.ReceiptDTO, tx *sql.Tx, receiptId int64, query string) error {
	for _, product := range receipt.Products {
		var itemId int64
		err := tx.QueryRow(db.PSCreateReceiptItem, product.Quantity, product.Id, product.Price).Scan(&itemId)
		if err != nil {
			return err
		}

		_, err = tx.Exec(query, receiptId, itemId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (receiptService *ReceiptService) GetInvoicesWithLimit(offset int, limit int) ([]model.AllInvoicesDTO, error) {
	rows, err := db.DBConn.Query(db.PSGetAllInvoices, limit, offset)
	if err != nil {
		return nil, err
	}

	var invoices []model.AllInvoicesDTO
	for rows.Next() {
		var invoice model.AllInvoicesDTO
		if err := rows.Scan(&invoice.Id, &invoice.CreatedAt, &invoice.ShopName, &invoice.CashierName, &invoice.CashierLastName, &invoice.EIN); err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

func (receiptService *ReceiptService) GetReceiptsWithLimit(offset int, limit int) ([]model.AllReceiptsDTO, error) {
	rows, err := db.DBConn.Query(db.PSGetAllReceipts, limit, offset)
	if err != nil {
		return nil, err
	}

	var receipts []model.AllReceiptsDTO
	for rows.Next() {
		var receipt model.AllReceiptsDTO
		if err := rows.Scan(&receipt.Id, &receipt.CreatedAt, &receipt.ShopName, &receipt.CashierName, &receipt.CashierLastName); err != nil {
			return nil, err
		}
		receipts = append(receipts, receipt)
	}

	return receipts, nil
}

func (receiptService *ReceiptService) CalculatePagesForAllReceipts(itemsPerPage int, query string) (model.AllReceiptsPages, error) {
	var count int

	err := db.DBConn.QueryRow(query).Scan(&count)
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
