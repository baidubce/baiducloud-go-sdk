package ocr

type TaxiReceiptWordsResult struct {
	InvoiceCode          *string `json:"InvoiceCode,omitempty"`
	InvoiceNum           *string `json:"InvoiceNum,omitempty"`
	TaxiNum              *string `json:"TaxiNum,omitempty"`
	Date                 *string `json:"Date,omitempty"`
	Time                 *string `json:"Time,omitempty"`
	PickupTime           *string `json:"PickupTime,omitempty"`
	DropoffTime          *string `json:"DropoffTime,omitempty"`
	Fare                 *string `json:"Fare,omitempty"`
	FuelOilSurcharge     *string `json:"FuelOilSurcharge,omitempty"`
	CallServiceSurcharge *string `json:"CallServiceSurcharge,omitempty"`
	TotalFare            *string `json:"TotalFare,omitempty"`
	Location             *string `json:"Location,omitempty"`
	Province             *string `json:"Province,omitempty"`
	City                 *string `json:"City,omitempty"`
	PricePerkm           *string `json:"PricePerkm,omitempty"`
	Distance             *string `json:"Distance,omitempty"`
	ServiceType          *string `json:"ServiceType,omitempty"`
}
