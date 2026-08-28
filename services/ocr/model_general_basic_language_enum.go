package ocr

// GeneralBasicLanguageEnum the model 'GeneralBasicLanguageEnum'
type GeneralBasicLanguageEnum int32

// List of GeneralBasicLanguageEnum
const (
	GeneralBasicLanguageEnumValueMinus1 GeneralBasicLanguageEnum = -1
	GeneralBasicLanguageEnumValue0      GeneralBasicLanguageEnum = 0
	GeneralBasicLanguageEnumValue1      GeneralBasicLanguageEnum = 1
	GeneralBasicLanguageEnumValue2      GeneralBasicLanguageEnum = 2
	GeneralBasicLanguageEnumValue3      GeneralBasicLanguageEnum = 3
)

// All allowed values of GeneralBasicLanguageEnum enum
var AllowedGeneralBasicLanguageEnumEnumValues = []GeneralBasicLanguageEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GeneralBasicLanguageEnum) IsValid() bool {
	for _, existing := range AllowedGeneralBasicLanguageEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
