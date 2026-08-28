package ocr

// IdCardCoverStatusEnum the model 'IdCardCoverStatusEnum'
type IdCardCoverStatusEnum int32

// List of IdCardCoverStatusEnum
const (
	IdCardCoverStatusEnumValue0 IdCardCoverStatusEnum = 0
	IdCardCoverStatusEnumValue1 IdCardCoverStatusEnum = 1
)

// All allowed values of IdCardCoverStatusEnum enum
var AllowedIdCardCoverStatusEnumEnumValues = []IdCardCoverStatusEnum{
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdCardCoverStatusEnum) IsValid() bool {
	for _, existing := range AllowedIdCardCoverStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
