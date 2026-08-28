package ocr

// WebImageLocDirectionEnum the model 'WebImageLocDirectionEnum'
type WebImageLocDirectionEnum int32

// List of WebImageLocDirectionEnum
const (
	WebImageLocDirectionEnumValueMinus1 WebImageLocDirectionEnum = -1
	WebImageLocDirectionEnumValue0      WebImageLocDirectionEnum = 0
	WebImageLocDirectionEnumValue1      WebImageLocDirectionEnum = 1
	WebImageLocDirectionEnumValue2      WebImageLocDirectionEnum = 2
	WebImageLocDirectionEnumValue3      WebImageLocDirectionEnum = 3
)

// All allowed values of WebImageLocDirectionEnum enum
var AllowedWebImageLocDirectionEnumEnumValues = []WebImageLocDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebImageLocDirectionEnum) IsValid() bool {
	for _, existing := range AllowedWebImageLocDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
