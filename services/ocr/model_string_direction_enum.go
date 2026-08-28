package ocr

// StringDirectionEnum the model 'StringDirectionEnum'
type StringDirectionEnum string

// List of StringDirectionEnum
const (
	StringDirectionEnumValueMinus1 StringDirectionEnum = "-1"
	StringDirectionEnumValue0      StringDirectionEnum = "0"
	StringDirectionEnumValue1      StringDirectionEnum = "1"
	StringDirectionEnumValue2      StringDirectionEnum = "2"
	StringDirectionEnumValue3      StringDirectionEnum = "3"
)

// All allowed values of StringDirectionEnum enum
var AllowedStringDirectionEnumEnumValues = []StringDirectionEnum{
	"-1",
	"0",
	"1",
	"2",
	"3",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StringDirectionEnum) IsValid() bool {
	for _, existing := range AllowedStringDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
