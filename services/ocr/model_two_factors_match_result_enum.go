package ocr

// TwoFactorsMatchResultEnum the model 'TwoFactorsMatchResultEnum'
type TwoFactorsMatchResultEnum string

// List of TwoFactorsMatchResultEnum
const (
	TwoFactorsMatchResultEnumValue0 TwoFactorsMatchResultEnum = "0"
	TwoFactorsMatchResultEnumValue1 TwoFactorsMatchResultEnum = "1"
	TwoFactorsMatchResultEnumValue2 TwoFactorsMatchResultEnum = "2"
)

// All allowed values of TwoFactorsMatchResultEnum enum
var AllowedTwoFactorsMatchResultEnumEnumValues = []TwoFactorsMatchResultEnum{
	"0",
	"1",
	"2",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TwoFactorsMatchResultEnum) IsValid() bool {
	for _, existing := range AllowedTwoFactorsMatchResultEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
