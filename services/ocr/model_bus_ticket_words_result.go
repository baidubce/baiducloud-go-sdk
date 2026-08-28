package ocr

type BusTicketWordsResult struct {
	InvoiceCode        *string `json:"InvoiceCode,omitempty"`
	InvoiceNum         *string `json:"InvoiceNum,omitempty"`
	Date               *string `json:"Date,omitempty"`
	Time               *string `json:"Time,omitempty"`
	StartingStation    *string `json:"StartingStation,omitempty"`
	Fare               *string `json:"Fare,omitempty"`
	IdNum              *string `json:"IdNum,omitempty"`
	DestinationStation *string `json:"DestinationStation,omitempty"`
	Name               *string `json:"Name,omitempty"`
}
