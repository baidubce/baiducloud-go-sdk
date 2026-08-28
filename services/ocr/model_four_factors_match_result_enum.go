package ocr

// FourFactorsMatchResultEnum the model 'FourFactorsMatchResultEnum'
type FourFactorsMatchResultEnum string

// List of FourFactorsMatchResultEnum
const (
	FourFactorsMatchResultEnumValue0 FourFactorsMatchResultEnum = "0"
	FourFactorsMatchResultEnumValue1 FourFactorsMatchResultEnum = "1"
	FourFactorsMatchResultEnumValue2 FourFactorsMatchResultEnum = "2"
)

// All allowed values of FourFactorsMatchResultEnum enum
var AllowedFourFactorsMatchResultEnumEnumValues = []FourFactorsMatchResultEnum{
	"0",
	"1",
	"2",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FourFactorsMatchResultEnum) IsValid() bool {
	for _, existing := range AllowedFourFactorsMatchResultEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
