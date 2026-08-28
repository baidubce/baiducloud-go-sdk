package ocr

// WebImageDirectionEnum the model 'WebImageDirectionEnum'
type WebImageDirectionEnum int32

// List of WebImageDirectionEnum
const (
	WebImageDirectionEnumValueMinus1 WebImageDirectionEnum = -1
	WebImageDirectionEnumValue0      WebImageDirectionEnum = 0
	WebImageDirectionEnumValue1      WebImageDirectionEnum = 1
	WebImageDirectionEnumValue2      WebImageDirectionEnum = 2
	WebImageDirectionEnumValue3      WebImageDirectionEnum = 3
)

// All allowed values of WebImageDirectionEnum enum
var AllowedWebImageDirectionEnumEnumValues = []WebImageDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebImageDirectionEnum) IsValid() bool {
	for _, existing := range AllowedWebImageDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
