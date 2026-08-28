package ocr

// MultiIdcardCompleteStatusEnum the model 'MultiIdcardCompleteStatusEnum'
type MultiIdcardCompleteStatusEnum int32

// List of MultiIdcardCompleteStatusEnum
const (
	MultiIdcardCompleteStatusEnumValue0 MultiIdcardCompleteStatusEnum = 0
	MultiIdcardCompleteStatusEnumValue1 MultiIdcardCompleteStatusEnum = 1
)

// All allowed values of MultiIdcardCompleteStatusEnum enum
var AllowedMultiIdcardCompleteStatusEnumEnumValues = []MultiIdcardCompleteStatusEnum{
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MultiIdcardCompleteStatusEnum) IsValid() bool {
	for _, existing := range AllowedMultiIdcardCompleteStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
