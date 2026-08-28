package ocr

// MultiIdcardNumberTypeEnum the model 'MultiIdcardNumberTypeEnum'
type MultiIdcardNumberTypeEnum int32

// List of MultiIdcardNumberTypeEnum
const (
	MultiIdcardNumberTypeEnumValueMinus1 MultiIdcardNumberTypeEnum = -1
	MultiIdcardNumberTypeEnumValue0      MultiIdcardNumberTypeEnum = 0
	MultiIdcardNumberTypeEnumValue1      MultiIdcardNumberTypeEnum = 1
	MultiIdcardNumberTypeEnumValue2      MultiIdcardNumberTypeEnum = 2
	MultiIdcardNumberTypeEnumValue3      MultiIdcardNumberTypeEnum = 3
	MultiIdcardNumberTypeEnumValue4      MultiIdcardNumberTypeEnum = 4
)

// All allowed values of MultiIdcardNumberTypeEnum enum
var AllowedMultiIdcardNumberTypeEnumEnumValues = []MultiIdcardNumberTypeEnum{
	-1,
	0,
	1,
	2,
	3,
	4,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MultiIdcardNumberTypeEnum) IsValid() bool {
	for _, existing := range AllowedMultiIdcardNumberTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
