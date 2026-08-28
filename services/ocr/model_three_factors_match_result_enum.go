package ocr

// ThreeFactorsMatchResultEnum the model 'ThreeFactorsMatchResultEnum'
type ThreeFactorsMatchResultEnum string

// List of ThreeFactorsMatchResultEnum
const (
	ThreeFactorsMatchResultEnumValue0 ThreeFactorsMatchResultEnum = "0"
	ThreeFactorsMatchResultEnumValue1 ThreeFactorsMatchResultEnum = "1"
	ThreeFactorsMatchResultEnumValue2 ThreeFactorsMatchResultEnum = "2"
)

// All allowed values of ThreeFactorsMatchResultEnum enum
var AllowedThreeFactorsMatchResultEnumEnumValues = []ThreeFactorsMatchResultEnum{
	"0",
	"1",
	"2",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ThreeFactorsMatchResultEnum) IsValid() bool {
	for _, existing := range AllowedThreeFactorsMatchResultEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
