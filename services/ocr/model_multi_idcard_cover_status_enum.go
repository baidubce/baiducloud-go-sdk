package ocr

// MultiIdcardCoverStatusEnum the model 'MultiIdcardCoverStatusEnum'
type MultiIdcardCoverStatusEnum int32

// List of MultiIdcardCoverStatusEnum
const (
	MultiIdcardCoverStatusEnumValue0 MultiIdcardCoverStatusEnum = 0
	MultiIdcardCoverStatusEnumValue1 MultiIdcardCoverStatusEnum = 1
)

// All allowed values of MultiIdcardCoverStatusEnum enum
var AllowedMultiIdcardCoverStatusEnumEnumValues = []MultiIdcardCoverStatusEnum{
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MultiIdcardCoverStatusEnum) IsValid() bool {
	for _, existing := range AllowedMultiIdcardCoverStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
