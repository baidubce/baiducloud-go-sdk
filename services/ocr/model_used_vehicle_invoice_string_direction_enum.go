package ocr

// UsedVehicleInvoiceStringDirectionEnum the model 'UsedVehicleInvoiceStringDirectionEnum'
type UsedVehicleInvoiceStringDirectionEnum string

// List of UsedVehicleInvoiceStringDirectionEnum
const (
	UsedVehicleInvoiceStringDirectionEnumValueMinus1 UsedVehicleInvoiceStringDirectionEnum = "-1"
	UsedVehicleInvoiceStringDirectionEnumValue0      UsedVehicleInvoiceStringDirectionEnum = "0"
	UsedVehicleInvoiceStringDirectionEnumValue1      UsedVehicleInvoiceStringDirectionEnum = "1"
	UsedVehicleInvoiceStringDirectionEnumValue2      UsedVehicleInvoiceStringDirectionEnum = "2"
	UsedVehicleInvoiceStringDirectionEnumValue3      UsedVehicleInvoiceStringDirectionEnum = "3"
)

// All allowed values of UsedVehicleInvoiceStringDirectionEnum enum
var AllowedUsedVehicleInvoiceStringDirectionEnumEnumValues = []UsedVehicleInvoiceStringDirectionEnum{
	"-1",
	"0",
	"1",
	"2",
	"3",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsedVehicleInvoiceStringDirectionEnum) IsValid() bool {
	for _, existing := range AllowedUsedVehicleInvoiceStringDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
