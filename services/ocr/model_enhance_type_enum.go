package ocr

// EnhanceTypeEnum the model 'EnhanceTypeEnum'
type EnhanceTypeEnum int32

// List of EnhanceTypeEnum
const (
	EnhanceTypeEnumValue0 EnhanceTypeEnum = 0
	EnhanceTypeEnumValue1 EnhanceTypeEnum = 1
	EnhanceTypeEnumValue2 EnhanceTypeEnum = 2
	EnhanceTypeEnumValue3 EnhanceTypeEnum = 3
)

// All allowed values of EnhanceTypeEnum enum
var AllowedEnhanceTypeEnumEnumValues = []EnhanceTypeEnum{
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnhanceTypeEnum) IsValid() bool {
	for _, existing := range AllowedEnhanceTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
