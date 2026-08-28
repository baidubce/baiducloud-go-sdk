package ocr

// HandwritingDirectionEnum the model 'HandwritingDirectionEnum'
type HandwritingDirectionEnum int32

// List of HandwritingDirectionEnum
const (
	HandwritingDirectionEnumValueMinus1 HandwritingDirectionEnum = -1
	HandwritingDirectionEnumValue0      HandwritingDirectionEnum = 0
	HandwritingDirectionEnumValue1      HandwritingDirectionEnum = 1
	HandwritingDirectionEnumValue2      HandwritingDirectionEnum = 2
	HandwritingDirectionEnumValue3      HandwritingDirectionEnum = 3
)

// All allowed values of HandwritingDirectionEnum enum
var AllowedHandwritingDirectionEnumEnumValues = []HandwritingDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HandwritingDirectionEnum) IsValid() bool {
	for _, existing := range AllowedHandwritingDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
