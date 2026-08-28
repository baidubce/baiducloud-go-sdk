package ocr

// ThreeFactorsVerifyResultEnum the model 'ThreeFactorsVerifyResultEnum'
type ThreeFactorsVerifyResultEnum string

// List of ThreeFactorsVerifyResultEnum
const (
	ThreeFactorsVerifyResultEnumValue1 ThreeFactorsVerifyResultEnum = "1"
	ThreeFactorsVerifyResultEnumValue0 ThreeFactorsVerifyResultEnum = "0"
	ThreeFactorsVerifyResultEnumValue2 ThreeFactorsVerifyResultEnum = "2"
)

// All allowed values of ThreeFactorsVerifyResultEnum enum
var AllowedThreeFactorsVerifyResultEnumEnumValues = []ThreeFactorsVerifyResultEnum{
	"1",
	"0",
	"2",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ThreeFactorsVerifyResultEnum) IsValid() bool {
	for _, existing := range AllowedThreeFactorsVerifyResultEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
