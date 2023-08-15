package converter

import "github.com/DzoniDiplomski/Backend_API/model"

func FromDTOReceipt(receiptDTO model.ReceiptDTO) model.Receipt {
	return model.Receipt{
		CashBoxId: receiptDTO.CashBoxId,
		ShopId:    receiptDTO.ShopId,
	}
}
