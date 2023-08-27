package service

import (
	"fmt"
	"strings"

	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/model"
	"github.com/DzoniDiplomski/Backend_API/repo"
)

type ProductService struct {
}

var productRepo = &repo.ProductRepo{}

func (productService *ProductService) AddProducts(products []model.Product) error {
	productIds, err := productRepo.AddProducts(products)
	if err != nil {
		return err
	}

	tx, err := db.DBConn.Begin()
	if err != nil {
		return err
	}

	for _, id := range productIds {
		_, err := tx.Exec(db.PSAddProductToStorage, id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (productService *ProductService) SearchForProducts(searchString string) ([]model.Product, error) {
	searchString = fmt.Sprintf("%%%s%%", searchString)
	rows, err := db.DBConn.Query(db.PSSearchProducts, searchString)
	if err != nil {
		return nil, err
	}

	var products []model.Product

	for rows.Next() {
		var product model.Product

		err := rows.Scan(&product.Id, &product.Name, &product.Barcode, &product.Quantity, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (productService *ProductService) UpdateProductPrice(price model.ProductDTO) error {
	tx, err := db.DBConn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	var priceId int64
	err = tx.QueryRow(db.PSUpdateProductPrice, price.Price, price.StartDate, price.EndDate).Scan(&priceId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(db.PSBindPriceWithProduct, priceId, price.Id, true)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(db.PSRevokeAllPrices, price.Id, priceId)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (productService *ProductService) GetProductPriceStats(id int64) ([]map[string]interface{}, error) {
	rows, err := db.DBConn.Query(db.PSGetProductPricesOverTime, id)
	if err != nil {
		return nil, err
	}

	var prices []map[string]interface{}
	for rows.Next() {
		var price float64
		var startDate, endDate string
		err := rows.Scan(&price, &startDate, &endDate)
		if err != nil {
			return nil, err
		}

		startDate = strings.Split(startDate, "T")[0]
		endDate = strings.Split(endDate, "T")[0]

		priceInterval := map[string]interface{}{
			"price":     price,
			"startDate": startDate,
			"endDate":   endDate,
		}

		prices = append(prices, priceInterval)
	}

	return prices, nil
}
