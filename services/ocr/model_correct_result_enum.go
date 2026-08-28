package ocr

// CorrectResultEnum the model 'CorrectResultEnum'
type CorrectResultEnum int32

// List of CorrectResultEnum
const (
	CorrectResultEnumValue0 CorrectResultEnum = 0
	CorrectResultEnumValue1 CorrectResultEnum = 1
	CorrectResultEnumValue2 CorrectResultEnum = 2
	CorrectResultEnumValue3 CorrectResultEnum = 3
)

// All allowed values of CorrectResultEnum enum
var AllowedCorrectResultEnumEnumValues = []CorrectResultEnum{
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CorrectResultEnum) IsValid() bool {
	for _, existing := range AllowedCorrectResultEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
