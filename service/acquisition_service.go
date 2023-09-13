package service

import (
	"fmt"
	"os"

	"github.com/DzoniDiplomski/Backend_API/model"
)

type AcquisitionService struct {
}

func (acquisitionService *AcquisitionService) GetAcquisitionNames(offset int, limit int) ([]string, error) {
	files, err := os.ReadDir("./acquisitions")
	if err != nil {
		return nil, err
	}

	fmt.Println(limit)
	var filenames []string
	if limit > len(files) {
		limit = len(files)
	}
	if offset == limit {
		limit++
	}
	files = files[offset:limit]
	for _, f := range files {
		filenames = append(filenames, f.Name())
	}
	return filenames, err
}

func (acquisitionService *AcquisitionService) OpenAcquisition(filename string) ([]byte, error) {
	pdfData, err := os.ReadFile("./acquisitions/" + filename)
	if err != nil {
		return nil, err
	}
	return pdfData, nil
}

func (acquisitionService *AcquisitionService) CalculatePagesForAllAcquisitions(itemsPerPage int) (model.AllReceiptsPages, error) {
	files, err := os.ReadDir("./acquisitions")
	if err != nil {
		return model.AllReceiptsPages{}, err
	}

	count := len(files)

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
