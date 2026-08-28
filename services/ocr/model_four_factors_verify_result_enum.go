package ocr

// FourFactorsVerifyResultEnum the model 'FourFactorsVerifyResultEnum'
type FourFactorsVerifyResultEnum string

// List of FourFactorsVerifyResultEnum
const (
	FourFactorsVerifyResultEnumValue1 FourFactorsVerifyResultEnum = "1"
	FourFactorsVerifyResultEnumValue0 FourFactorsVerifyResultEnum = "0"
	FourFactorsVerifyResultEnumValue2 FourFactorsVerifyResultEnum = "2"
)

// All allowed values of FourFactorsVerifyResultEnum enum
var AllowedFourFactorsVerifyResultEnumEnumValues = []FourFactorsVerifyResultEnum{
	"1",
	"0",
	"2",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FourFactorsVerifyResultEnum) IsValid() bool {
	for _, existing := range AllowedFourFactorsVerifyResultEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
