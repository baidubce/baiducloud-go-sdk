package ocr

// GeneralLanguageEnum the model 'GeneralLanguageEnum'
type GeneralLanguageEnum int32

// List of GeneralLanguageEnum
const (
	GeneralLanguageEnumValueMinus1 GeneralLanguageEnum = -1
	GeneralLanguageEnumValue0      GeneralLanguageEnum = 0
	GeneralLanguageEnumValue1      GeneralLanguageEnum = 1
	GeneralLanguageEnumValue2      GeneralLanguageEnum = 2
	GeneralLanguageEnumValue3      GeneralLanguageEnum = 3
)

// All allowed values of GeneralLanguageEnum enum
var AllowedGeneralLanguageEnumEnumValues = []GeneralLanguageEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GeneralLanguageEnum) IsValid() bool {
	for _, existing := range AllowedGeneralLanguageEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
