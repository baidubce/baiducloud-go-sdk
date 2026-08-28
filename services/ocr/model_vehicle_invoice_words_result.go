package ocr

type VehicleInvoiceWordsResult struct {
	InvoiceHeader        *string `json:"InvoiceHeader,omitempty"`
	InvoiceCode          *string `json:"InvoiceCode,omitempty"`
	InvoiceNum           *string `json:"InvoiceNum,omitempty"`
	PrintedCode          *string `json:"PrintedCode,omitempty"`
	PrintedNum           *string `json:"PrintedNum,omitempty"`
	InvoiceDate          *string `json:"InvoiceDate,omitempty"`
	MachineCode          *string `json:"MachineCode,omitempty"`
	Purchaser            *string `json:"Purchaser,omitempty"`
	PurchaserCode        *string `json:"PurchaserCode,omitempty"`
	VehicleType          *string `json:"VehicleType,omitempty"`
	ManuModel            *string `json:"ManuModel,omitempty"`
	Origin               *string `json:"Origin,omitempty"`
	CertificateNum       *string `json:"CertificateNum,omitempty"`
	EngineNum            *string `json:"EngineNum,omitempty"`
	VinNum               *string `json:"VinNum,omitempty"`
	PriceTax             *string `json:"PriceTax,omitempty"`
	PriceTaxLow          *string `json:"PriceTaxLow,omitempty"`
	Saler                *string `json:"Saler,omitempty"`
	SalerPhone           *string `json:"SalerPhone,omitempty"`
	SalerCode            *string `json:"SalerCode,omitempty"`
	SalerAccountNum      *string `json:"SalerAccountNum,omitempty"`
	SalerAddress         *string `json:"SalerAddress,omitempty"`
	SalerBank            *string `json:"SalerBank,omitempty"`
	TaxRate              *string `json:"TaxRate,omitempty"`
	Tax                  *string `json:"Tax,omitempty"`
	TaxAuthor            *string `json:"TaxAuthor,omitempty"`
	TaxAuthorCode        *string `json:"TaxAuthorCode,omitempty"`
	Price                *string `json:"Price,omitempty"`
	LimitPassenger       *string `json:"LimitPassenger,omitempty"`
	Toonage              *string `json:"toonage,omitempty"`
	SheetNum             *string `json:"sheet-num,omitempty"`
	Drawer               *string `json:"drawer,omitempty"`
	Remarks              *string `json:"remarks,omitempty"`
	ImportCertificateNum *string `json:"import-certificate-num,omitempty"`
	TaxPaymentVoucherNo  *string `json:"tax-payment-voucher-no,omitempty"`
	InspectionFormNum    *string `json:"inspection-form-num,omitempty"`
	TaxCode              *string `json:"tax-code,omitempty"`
	InvoiceNumDigit      *string `json:"InvoiceNumDigit,omitempty"`
}
