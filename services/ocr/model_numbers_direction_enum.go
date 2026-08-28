package ocr

// NumbersDirectionEnum the model 'NumbersDirectionEnum'
type NumbersDirectionEnum int32

// List of NumbersDirectionEnum
const (
	NumbersDirectionEnumValueMinus1 NumbersDirectionEnum = -1
	NumbersDirectionEnumValue0      NumbersDirectionEnum = 0
	NumbersDirectionEnumValue1      NumbersDirectionEnum = 1
	NumbersDirectionEnumValue2      NumbersDirectionEnum = 2
	NumbersDirectionEnumValue3      NumbersDirectionEnum = 3
)

// All allowed values of NumbersDirectionEnum enum
var AllowedNumbersDirectionEnumEnumValues = []NumbersDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NumbersDirectionEnum) IsValid() bool {
	for _, existing := range AllowedNumbersDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
