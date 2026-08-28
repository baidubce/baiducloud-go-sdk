package ocr

// IdCardCompleteStatusEnum the model 'IdCardCompleteStatusEnum'
type IdCardCompleteStatusEnum int32

// List of IdCardCompleteStatusEnum
const (
	IdCardCompleteStatusEnumValue0 IdCardCompleteStatusEnum = 0
	IdCardCompleteStatusEnumValue1 IdCardCompleteStatusEnum = 1
)

// All allowed values of IdCardCompleteStatusEnum enum
var AllowedIdCardCompleteStatusEnumEnumValues = []IdCardCompleteStatusEnum{
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdCardCompleteStatusEnum) IsValid() bool {
	for _, existing := range AllowedIdCardCompleteStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
