package ocr

// GeneralBasicDirectionEnum the model 'GeneralBasicDirectionEnum'
type GeneralBasicDirectionEnum int32

// List of GeneralBasicDirectionEnum
const (
	GeneralBasicDirectionEnumValueMinus1 GeneralBasicDirectionEnum = -1
	GeneralBasicDirectionEnumValue0      GeneralBasicDirectionEnum = 0
	GeneralBasicDirectionEnumValue1      GeneralBasicDirectionEnum = 1
	GeneralBasicDirectionEnumValue2      GeneralBasicDirectionEnum = 2
	GeneralBasicDirectionEnumValue3      GeneralBasicDirectionEnum = 3
)

// All allowed values of GeneralBasicDirectionEnum enum
var AllowedGeneralBasicDirectionEnumEnumValues = []GeneralBasicDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GeneralBasicDirectionEnum) IsValid() bool {
	for _, existing := range AllowedGeneralBasicDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
