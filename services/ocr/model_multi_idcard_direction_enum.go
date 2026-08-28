package ocr

// MultiIdcardDirectionEnum the model 'MultiIdcardDirectionEnum'
type MultiIdcardDirectionEnum int32

// List of MultiIdcardDirectionEnum
const (
	MultiIdcardDirectionEnumValueMinus1 MultiIdcardDirectionEnum = -1
	MultiIdcardDirectionEnumValue0      MultiIdcardDirectionEnum = 0
	MultiIdcardDirectionEnumValue1      MultiIdcardDirectionEnum = 1
	MultiIdcardDirectionEnumValue2      MultiIdcardDirectionEnum = 2
	MultiIdcardDirectionEnumValue3      MultiIdcardDirectionEnum = 3
)

// All allowed values of MultiIdcardDirectionEnum enum
var AllowedMultiIdcardDirectionEnumEnumValues = []MultiIdcardDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MultiIdcardDirectionEnum) IsValid() bool {
	for _, existing := range AllowedMultiIdcardDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
