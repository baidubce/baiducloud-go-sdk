package ocr

// FinanceInvoiceDirectionEnum the model 'FinanceInvoiceDirectionEnum'
type FinanceInvoiceDirectionEnum int32

// List of FinanceInvoiceDirectionEnum
const (
	FinanceInvoiceDirectionEnumValueMinus1 FinanceInvoiceDirectionEnum = -1
	FinanceInvoiceDirectionEnumValue0      FinanceInvoiceDirectionEnum = 0
	FinanceInvoiceDirectionEnumValue1      FinanceInvoiceDirectionEnum = 1
	FinanceInvoiceDirectionEnumValue2      FinanceInvoiceDirectionEnum = 2
	FinanceInvoiceDirectionEnumValue3      FinanceInvoiceDirectionEnum = 3
)

// All allowed values of FinanceInvoiceDirectionEnum enum
var AllowedFinanceInvoiceDirectionEnumEnumValues = []FinanceInvoiceDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FinanceInvoiceDirectionEnum) IsValid() bool {
	for _, existing := range AllowedFinanceInvoiceDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
