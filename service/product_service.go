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

func (productService *ProductService) AddProducts(products []model.Product) error {
	if err := productRepo.AddProducts(products); err != nil {
		return err
	}
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

		err := rows.Scan(&product.Id, &product.Barcode, &product.Name)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
