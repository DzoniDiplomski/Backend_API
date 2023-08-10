package service

import (
	"fmt"

	"diplomski.com/db"
	"diplomski.com/model"
	"diplomski.com/repo"
)

type ProductService struct {
}

var productRepo = &repo.ProductRepo{}

func (productService *ProductService) AddProducts(products []model.ProductFromSearch) error {
	if err := productRepo.AddProducts(products); err != nil {
		return err
	}
	return nil
}

func (productService *ProductService) SearchForProducts(searchString string) ([]model.ProductFromSearch, error) {
	searchString = fmt.Sprintf("%%%s%%", searchString)
	rows, err := db.DBConn.Query(db.PSSearchProducts, searchString)
	if err != nil {
		return nil, err
	}

	var products []model.ProductFromSearch

	for rows.Next() {
		var product model.ProductFromSearch

		err := rows.Scan(&product.Id, &product.Name, &product.Barcode, &product.Quantity, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
