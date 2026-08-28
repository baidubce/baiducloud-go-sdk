package ocr

// IdCardDirectionEnum the model 'IdCardDirectionEnum'
type IdCardDirectionEnum int32

// List of IdCardDirectionEnum
const (
	IdCardDirectionEnumValueMinus1 IdCardDirectionEnum = -1
	IdCardDirectionEnumValue0      IdCardDirectionEnum = 0
	IdCardDirectionEnumValue1      IdCardDirectionEnum = 1
	IdCardDirectionEnumValue2      IdCardDirectionEnum = 2
	IdCardDirectionEnumValue3      IdCardDirectionEnum = 3
)

// All allowed values of IdCardDirectionEnum enum
var AllowedIdCardDirectionEnumEnumValues = []IdCardDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdCardDirectionEnum) IsValid() bool {
	for _, existing := range AllowedIdCardDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
