package service

import (
	"database/sql"
	"fmt"

	"github.com/DzoniDiplomski/Backend_API/converter"
	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/model"
	"github.com/DzoniDiplomski/Backend_API/repo"
)

type ReceiptService struct {
}

var receiptRepo = &repo.ReceiptRepo{}

func (receiptService *ReceiptService) GetReceiptItems(receiptId int64) ([]model.Product, error) {
	rows, err := db.DBConn.Query(db.PSGetReceiptItems, receiptId)
	if err != nil {
		return nil, err
	}

	var (
		product model.Product
		items   []model.Product
	)
	for rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Quantity, &product.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, product)
	}
	return items, nil
}

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
		addSumQuery      string
	)

	if receipt.EIN != 0 {
		createQuery = db.PSAddInvoice
		bindItemQuery = db.PSBindItemWithInvoice
		bindCashierQuery = db.PSBindInvoiceWithCashier
		deleteQuery = db.PSDeleteInvoice
		addSumQuery = db.PSAddSumToInvoice
	} else {
		createQuery = db.PSAddReceipt
		bindItemQuery = db.PSBindItemWithReceipt
		bindCashierQuery = db.PSBindReceiptWithCashier
		deleteQuery = db.PSDeleteReceipt
		addSumQuery = db.PSAddSumToReceipt
	}

	receiptId, err := receiptRepo.Create(converter.FromDTOReceipt(receipt), createQuery)
	if err != nil {
		tx.Rollback()
		return err
	}

	err, sum := addItemsToDBAndBindWithReceipt(receipt, tx, receiptId, bindItemQuery)
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

	err = addSumToReceipt(tx, receiptId, sum, addSumQuery)
	if err != nil {
		receiptRepo.Delete(receiptId, deleteQuery)
		tx.Rollback()
		return err
	}

	err = addReceiptSumToDailyMarket(tx, sum)
	if err != nil {
		receiptRepo.Delete(receiptId, deleteQuery)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func bindReceiptWithCashier(tx *sql.Tx, receiptId int64, receipt model.ReceiptDTO, query string) error {
	args := []interface{}{receiptId, receipt.CashierId}
	if query == db.PSBindInvoiceWithCashier {
		args = []interface{}{receiptId, receipt.CashierId, receipt.EIN}
	}
	_, err := tx.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func addSumToReceipt(tx *sql.Tx, receiptId int64, sum float64, query string) error {
	_, err := tx.Exec(query, sum, receiptId)
	if err != nil {
		return err
	}
	return nil
}

func addReceiptSumToDailyMarket(tx *sql.Tx, sum float64) error {
	var dailyMarket model.DailyMarket
	err := tx.QueryRow(db.PSGetTodaysMarket).Scan(&dailyMarket.Id, &dailyMarket.Date, &dailyMarket.Sum)
	if err == sql.ErrNoRows {
		fmt.Println("Creating a new market for today!")
		err := tx.QueryRow(db.PSCreateTodaysMarket).Scan(&dailyMarket.Id)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	} else if err != nil {
		return err
	}

	_, err = tx.Exec(db.PSUpdateTodaysMarketSum, sum, dailyMarket.Id)
	if err != nil {
		return err
	}
	return nil
}

func addItemsToDBAndBindWithReceipt(receipt model.ReceiptDTO, tx *sql.Tx, receiptId int64, query string) (error, float64) {
	sum := 0.0
	for _, product := range receipt.Products {
		var itemId int64
		err := tx.QueryRow(db.PSCreateReceiptItem, product.Quantity, product.Id, product.Price).Scan(&itemId)
		if err != nil {
			return err, 0.0
		}

		_, err = tx.Exec(query, receiptId, itemId)
		if err != nil {
			return err, 0.0
		}

		err = reduceItemQuantityInDB(product.Quantity, product.Id, tx)
		if err != nil {
			return err, 0.0
		}

		sum += float64(product.Quantity) * product.Price
	}
	return nil, sum
}

func reduceItemQuantityInDB(quantity, itemId int, tx *sql.Tx) error {
	_, err := tx.Exec(db.PSReduceItemQuantity, quantity, itemId)
	if err != nil {
		return err
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
			fmt.Println(receipt)
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
