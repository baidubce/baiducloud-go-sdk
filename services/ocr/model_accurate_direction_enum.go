package ocr

// AccurateDirectionEnum the model 'AccurateDirectionEnum'
type AccurateDirectionEnum int32

// List of AccurateDirectionEnum
const (
	AccurateDirectionEnumValueMinus1 AccurateDirectionEnum = -1
	AccurateDirectionEnumValue0      AccurateDirectionEnum = 0
	AccurateDirectionEnumValue1      AccurateDirectionEnum = 1
	AccurateDirectionEnumValue2      AccurateDirectionEnum = 2
	AccurateDirectionEnumValue3      AccurateDirectionEnum = 3
)

// All allowed values of AccurateDirectionEnum enum
var AllowedAccurateDirectionEnumEnumValues = []AccurateDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccurateDirectionEnum) IsValid() bool {
	for _, existing := range AllowedAccurateDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
