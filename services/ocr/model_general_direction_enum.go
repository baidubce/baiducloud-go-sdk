package ocr

// GeneralDirectionEnum the model 'GeneralDirectionEnum'
type GeneralDirectionEnum int32

// List of GeneralDirectionEnum
const (
	GeneralDirectionEnumValueMinus1 GeneralDirectionEnum = -1
	GeneralDirectionEnumValue0      GeneralDirectionEnum = 0
	GeneralDirectionEnumValue1      GeneralDirectionEnum = 1
	GeneralDirectionEnumValue2      GeneralDirectionEnum = 2
	GeneralDirectionEnumValue3      GeneralDirectionEnum = 3
)

// All allowed values of GeneralDirectionEnum enum
var AllowedGeneralDirectionEnumEnumValues = []GeneralDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GeneralDirectionEnum) IsValid() bool {
	for _, existing := range AllowedGeneralDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
