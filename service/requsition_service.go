package service

import (
	"github.com/DzoniDiplomski/Backend_API/converter"
	"github.com/DzoniDiplomski/Backend_API/model"
	"github.com/DzoniDiplomski/Backend_API/repo"
)

type RequisitionService struct {
}

var requisitionRepo = repo.RequisitionRepo{}

func (requisitionService *RequisitionService) CreateRequisition(requisition model.RequisitionDTO) error {
	id, err := requisitionRepo.Create(converter.FromDTORequisition(requisition))
	if err != nil {
		return err
	}

}

func (requisitionService *RequisitionService) createRequisitionItem(products []model.Product) {
	for _, product := range products {

	}
}
