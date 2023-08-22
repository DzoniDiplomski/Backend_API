package converter

import "github.com/DzoniDiplomski/Backend_API/model"

func FromDTORequisition(requisitionDTO model.RequisitionDTO) model.Requisition {
	return model.Requisition{
		ManagerId: requisitionDTO.ManagerId,
	}
}
