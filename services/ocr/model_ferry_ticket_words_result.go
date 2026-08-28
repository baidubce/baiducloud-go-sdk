package ocr

type FerryTicketWordsResult struct {
	InvoiceType        *string `json:"InvoiceType,omitempty"`
	InvoiceCode        *string `json:"InvoiceCode,omitempty"`
	InvoiceNum         *string `json:"InvoiceNum,omitempty"`
	StartingStation    *string `json:"StartingStation,omitempty"`
	DestinationStation *string `json:"DestinationStation,omitempty"`
	Fare               *string `json:"Fare,omitempty"`
	InvoiceDate        *string `json:"InvoiceDate,omitempty"`
	BarCode            *string `json:"BarCode,omitempty"`
	BarCodeNum         *string `json:"BarCodeNum,omitempty"`
	City               *string `json:"City,omitempty"`
	InvoiceTitle       *string `json:"InvoiceTitle,omitempty"`
	Province           *string `json:"Province,omitempty"`
	QrCode             *string `json:"QrCode,omitempty"`
	Time               *string `json:"Time,omitempty"`
	TicketTime         *string `json:"TicketTime,omitempty"`
	TicketDate         *string `json:"TicketDate,omitempty"`
	IdCard             *string `json:"IdCard,omitempty"`
	PassengerName      *string `json:"PassengerName,omitempty"`
}
