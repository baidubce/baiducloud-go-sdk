package ocr

// TwoFactorsVerifyResultEnum the model 'TwoFactorsVerifyResultEnum'
type TwoFactorsVerifyResultEnum string

// List of TwoFactorsVerifyResultEnum
const (
	TwoFactorsVerifyResultEnumValue1 TwoFactorsVerifyResultEnum = "1"
	TwoFactorsVerifyResultEnumValue0 TwoFactorsVerifyResultEnum = "0"
	TwoFactorsVerifyResultEnumValue2 TwoFactorsVerifyResultEnum = "2"
)

// All allowed values of TwoFactorsVerifyResultEnum enum
var AllowedTwoFactorsVerifyResultEnumEnumValues = []TwoFactorsVerifyResultEnum{
	"1",
	"0",
	"2",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TwoFactorsVerifyResultEnum) IsValid() bool {
	for _, existing := range AllowedTwoFactorsVerifyResultEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
