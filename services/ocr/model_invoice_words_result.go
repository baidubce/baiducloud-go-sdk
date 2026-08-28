package ocr

type InvoiceWordsResult struct {
	InvoiceType          *string           `json:"InvoiceType,omitempty"`
	InvoiceCode          *string           `json:"InvoiceCode,omitempty"`
	InvoiceNum           *string           `json:"InvoiceNum,omitempty"`
	InvoiceDate          *string           `json:"InvoiceDate,omitempty"`
	AmountInFiguers      *string           `json:"AmountInFiguers,omitempty"`
	AmountInWords        *string           `json:"AmountInWords,omitempty"`
	CommodityName        []*InvoiceRowWord `json:"CommodityName,omitempty"`
	CommodityUnit        []*InvoiceRowWord `json:"CommodityUnit,omitempty"`
	CommodityPrice       []*InvoiceRowWord `json:"CommodityPrice,omitempty"`
	CommodityNum         []*InvoiceRowWord `json:"CommodityNum,omitempty"`
	CommodityAmount      []*InvoiceRowWord `json:"CommodityAmount,omitempty"`
	IndustrySort         *string           `json:"IndustrySort,omitempty"`
	MachineNum           *string           `json:"MachineNum,omitempty"`
	CheckCode            *string           `json:"CheckCode,omitempty"`
	SellerName           *string           `json:"SellerName,omitempty"`
	SellerRegisterNum    *string           `json:"SellerRegisterNum,omitempty"`
	PurchaserName        *string           `json:"PurchaserName,omitempty"`
	PurchaserRegisterNum *string           `json:"PurchaserRegisterNum,omitempty"`
	TotalTax             *string           `json:"TotalTax,omitempty"`
	Province             *string           `json:"Province,omitempty"`
	City                 *string           `json:"City,omitempty"`
	Time                 *string           `json:"Time,omitempty"`
	SheetNum             *string           `json:"SheetNum,omitempty"`
}
