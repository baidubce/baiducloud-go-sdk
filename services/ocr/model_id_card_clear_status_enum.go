package ocr

// IdCardClearStatusEnum the model 'IdCardClearStatusEnum'
type IdCardClearStatusEnum int32

// List of IdCardClearStatusEnum
const (
	IdCardClearStatusEnumValue0 IdCardClearStatusEnum = 0
	IdCardClearStatusEnumValue1 IdCardClearStatusEnum = 1
)

// All allowed values of IdCardClearStatusEnum enum
var AllowedIdCardClearStatusEnumEnumValues = []IdCardClearStatusEnum{
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdCardClearStatusEnum) IsValid() bool {
	for _, existing := range AllowedIdCardClearStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
