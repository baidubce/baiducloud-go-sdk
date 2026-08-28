package ocr

type TrainTicketWordsResult struct {
	TicketNum          *string `json:"ticket_num,omitempty"`
	StartingStation    *string `json:"starting_station,omitempty"`
	TrainNum           *string `json:"train_num,omitempty"`
	DestinationStation *string `json:"destination_station,omitempty"`
	Date               *string `json:"date,omitempty"`
	TicketRates        *string `json:"ticket_rates,omitempty"`
	SeatCategory       *string `json:"seat_category,omitempty"`
	Name               *string `json:"name,omitempty"`
	IdNum              *string `json:"id_num,omitempty"`
	SerialNumber       *string `json:"serial_number,omitempty"`
	SalesStation       *string `json:"sales_station,omitempty"`
	Time               *string `json:"time,omitempty"`
	SeatNum            *string `json:"seat_num,omitempty"`
	RefundFlag         *string `json:"refund_flag,omitempty"`
	InvoiceNum         *string `json:"invoice_num,omitempty"`
	InvoiceDate        *string `json:"invoice_date,omitempty"`
	Fare               *string `json:"fare,omitempty"`
	TaxRate            *string `json:"tax_rate,omitempty"`
	Tax                *string `json:"tax,omitempty"`
	ElecTicketNum      *string `json:"elec_ticket_num,omitempty"`
	ServiceType        *string `json:"ServiceType,omitempty"`
}
