package ocr

// IsOnStockEnum the model 'IsOnStockEnum'
type IsOnStockEnum string

// List of IsOnStockEnum
const (
	IsOnStockEnumValue0 IsOnStockEnum = "0"
	IsOnStockEnumValue1 IsOnStockEnum = "1"
)

// All allowed values of IsOnStockEnum enum
var AllowedIsOnStockEnumEnumValues = []IsOnStockEnum{
	"0",
	"1",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IsOnStockEnum) IsValid() bool {
	for _, existing := range AllowedIsOnStockEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
