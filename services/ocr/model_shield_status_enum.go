package ocr

// ShieldStatusEnum the model 'ShieldStatusEnum'
type ShieldStatusEnum string

// List of ShieldStatusEnum
const (
	ShieldStatusEnumValue0 ShieldStatusEnum = "0"
	ShieldStatusEnumValue1 ShieldStatusEnum = "1"
)

// All allowed values of ShieldStatusEnum enum
var AllowedShieldStatusEnumEnumValues = []ShieldStatusEnum{
	"0",
	"1",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShieldStatusEnum) IsValid() bool {
	for _, existing := range AllowedShieldStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
