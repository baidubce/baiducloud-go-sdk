package ocr

// VatInvoiceTypeEnum the model 'VatInvoiceTypeEnum'
type VatInvoiceTypeEnum string

// List of VatInvoiceTypeEnum
const (
	VatInvoiceTypeEnumNormal VatInvoiceTypeEnum = "normal"
	VatInvoiceTypeEnumRoll   VatInvoiceTypeEnum = "roll"
)

// All allowed values of VatInvoiceTypeEnum enum
var AllowedVatInvoiceTypeEnumEnumValues = []VatInvoiceTypeEnum{
	"normal",
	"roll",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VatInvoiceTypeEnum) IsValid() bool {
	for _, existing := range AllowedVatInvoiceTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
