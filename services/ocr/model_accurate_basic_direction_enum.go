package ocr

// AccurateBasicDirectionEnum the model 'AccurateBasicDirectionEnum'
type AccurateBasicDirectionEnum int32

// List of AccurateBasicDirectionEnum
const (
	AccurateBasicDirectionEnumValueMinus1 AccurateBasicDirectionEnum = -1
	AccurateBasicDirectionEnumValue0      AccurateBasicDirectionEnum = 0
	AccurateBasicDirectionEnumValue1      AccurateBasicDirectionEnum = 1
	AccurateBasicDirectionEnumValue2      AccurateBasicDirectionEnum = 2
	AccurateBasicDirectionEnumValue3      AccurateBasicDirectionEnum = 3
)

// All allowed values of AccurateBasicDirectionEnum enum
var AllowedAccurateBasicDirectionEnumEnumValues = []AccurateBasicDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccurateBasicDirectionEnum) IsValid() bool {
	for _, existing := range AllowedAccurateBasicDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
