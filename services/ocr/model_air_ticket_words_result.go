package ocr

type AirTicketWordsResult struct {
	InvoiceTypeOrg       *string              `json:"invoice_type_org,omitempty"`
	Name                 []*AirTicketWordItem `json:"name,omitempty"`
	StartingStation      []*AirTicketWordItem `json:"starting_station,omitempty"`
	DestinationStation   []*AirTicketWordItem `json:"destination_station,omitempty"`
	Flight               []*AirTicketWordItem `json:"flight,omitempty"`
	Date                 []*AirTicketWordItem `json:"date,omitempty"`
	TicketNumber         []*AirTicketWordItem `json:"ticket_number,omitempty"`
	Fare                 []*AirTicketWordItem `json:"fare,omitempty"`
	DevFund              []*AirTicketWordItem `json:"dev_fund,omitempty"`
	FuelSurcharge        []*AirTicketWordItem `json:"fuel_surcharge,omitempty"`
	OtherTax             []*AirTicketWordItem `json:"other_tax,omitempty"`
	TicketRates          []*AirTicketWordItem `json:"ticket_rates,omitempty"`
	IssuedDate           []*AirTicketWordItem `json:"issued_date,omitempty"`
	IdNum                []*AirTicketWordItem `json:"id_num,omitempty"`
	Carrier              []*AirTicketWordItem `json:"carrier,omitempty"`
	Time                 []*AirTicketWordItem `json:"time,omitempty"`
	IssuedBy             []*AirTicketWordItem `json:"issued_by,omitempty"`
	SerialNumber         []*AirTicketWordItem `json:"serial_number,omitempty"`
	Insurance            []*AirTicketWordItem `json:"insurance,omitempty"`
	FareBasis            []*AirTicketWordItem `json:"fare_basis,omitempty"`
	Class                []*AirTicketWordItem `json:"class,omitempty"`
	AgentCode            []*AirTicketWordItem `json:"agent_code,omitempty"`
	Endorsement          []*AirTicketWordItem `json:"endorsement,omitempty"`
	Allow                []*AirTicketWordItem `json:"allow,omitempty"`
	Ck                   []*AirTicketWordItem `json:"ck,omitempty"`
	EffectiveDate        []*AirTicketWordItem `json:"effective_date,omitempty"`
	ExpirationDate       []*AirTicketWordItem `json:"expiration_date,omitempty"`
	InvoiceNum           []*AirTicketWordItem `json:"invoice_num,omitempty"`
	CommodityTaxRate     []*AirTicketWordItem `json:"commodity_tax_rate,omitempty"`
	CommodityTax         []*AirTicketWordItem `json:"commodity_tax,omitempty"`
	PurchaserName        []*AirTicketWordItem `json:"purchaser_name,omitempty"`
	PurchaserRegisterNum []*AirTicketWordItem `json:"purchaser_register_num,omitempty"`
	Identification       []*AirTicketWordItem `json:"identification,omitempty"`
	InvoiceStatus        []*AirTicketWordItem `json:"invoice_status,omitempty"`
	Tip                  []*AirTicketWordItem `json:"tip,omitempty"`
	ServiceType          []*AirTicketWordItem `json:"ServiceType,omitempty"`
}
