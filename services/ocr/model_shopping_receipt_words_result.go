package ocr

type ShoppingReceiptWordsResult struct {
	ShopName        *string  `json:"shop_name,omitempty"`
	ReceiptNum      *string  `json:"receipt_num,omitempty"`
	MachineNum      *string  `json:"machine_num,omitempty"`
	EmployeeNum     *string  `json:"employee_num,omitempty"`
	ConsumptionDate *string  `json:"consumption_date,omitempty"`
	ConsumptionTime *string  `json:"consumption_time,omitempty"`
	TotalAmount     *string  `json:"total_amount,omitempty"`
	Change          *string  `json:"change,omitempty"`
	Currency        *string  `json:"currency,omitempty"`
	PaidAmount      *string  `json:"paid_amount,omitempty"`
	Discount        *string  `json:"discount,omitempty"`
	PrintDate       *string  `json:"print_date,omitempty"`
	PrintTime       *string  `json:"print_time,omitempty"`
	TableRowNum     *int32   `json:"table_row_num,omitempty"`
	Table           []*Table `json:"table,omitempty"`
}
