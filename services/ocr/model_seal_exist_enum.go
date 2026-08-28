package ocr

// SealExistEnum the model 'SealExistEnum'
type SealExistEnum string

// List of SealExistEnum
const (
	SealExistEnumValue0 SealExistEnum = "0"
	SealExistEnumValue1 SealExistEnum = "1"
)

// All allowed values of SealExistEnum enum
var AllowedSealExistEnumEnumValues = []SealExistEnum{
	"0",
	"1",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SealExistEnum) IsValid() bool {
	for _, existing := range AllowedSealExistEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
