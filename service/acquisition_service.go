package service

import "os"

type AcquisitionService struct {
}

func (acquisitionService *AcquisitionService) GetAcquisitionNames() ([]string, error) {
	files, err := os.ReadDir("./acquisitions")
	if err != nil {
		return nil, err
	}

	var filenames []string
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
