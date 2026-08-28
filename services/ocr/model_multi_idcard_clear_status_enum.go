package ocr

// MultiIdcardClearStatusEnum the model 'MultiIdcardClearStatusEnum'
type MultiIdcardClearStatusEnum int32

// List of MultiIdcardClearStatusEnum
const (
	MultiIdcardClearStatusEnumValue0 MultiIdcardClearStatusEnum = 0
	MultiIdcardClearStatusEnumValue1 MultiIdcardClearStatusEnum = 1
)

// All allowed values of MultiIdcardClearStatusEnum enum
var AllowedMultiIdcardClearStatusEnumEnumValues = []MultiIdcardClearStatusEnum{
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MultiIdcardClearStatusEnum) IsValid() bool {
	for _, existing := range AllowedMultiIdcardClearStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
