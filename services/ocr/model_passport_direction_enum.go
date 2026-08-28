package ocr

// PassportDirectionEnum the model 'PassportDirectionEnum'
type PassportDirectionEnum int32

// List of PassportDirectionEnum
const (
	PassportDirectionEnumValueMinus1 PassportDirectionEnum = -1
	PassportDirectionEnumValue0      PassportDirectionEnum = 0
	PassportDirectionEnumValue1      PassportDirectionEnum = 1
	PassportDirectionEnumValue2      PassportDirectionEnum = 2
	PassportDirectionEnumValue3      PassportDirectionEnum = 3
)

// All allowed values of PassportDirectionEnum enum
var AllowedPassportDirectionEnumEnumValues = []PassportDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PassportDirectionEnum) IsValid() bool {
	for _, existing := range AllowedPassportDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
