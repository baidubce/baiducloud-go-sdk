package ocr

// IdCardNumberTypeEnum the model 'IdCardNumberTypeEnum'
type IdCardNumberTypeEnum int32

// List of IdCardNumberTypeEnum
const (
	IdCardNumberTypeEnumValueMinus1 IdCardNumberTypeEnum = -1
	IdCardNumberTypeEnumValue0      IdCardNumberTypeEnum = 0
	IdCardNumberTypeEnumValue1      IdCardNumberTypeEnum = 1
	IdCardNumberTypeEnumValue2      IdCardNumberTypeEnum = 2
	IdCardNumberTypeEnumValue3      IdCardNumberTypeEnum = 3
	IdCardNumberTypeEnumValue4      IdCardNumberTypeEnum = 4
)

// All allowed values of IdCardNumberTypeEnum enum
var AllowedIdCardNumberTypeEnumEnumValues = []IdCardNumberTypeEnum{
	-1,
	0,
	1,
	2,
	3,
	4,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdCardNumberTypeEnum) IsValid() bool {
	for _, existing := range AllowedIdCardNumberTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
