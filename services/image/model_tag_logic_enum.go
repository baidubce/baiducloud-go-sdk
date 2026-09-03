package image

// TagLogicEnum the model 'TagLogicEnum'
type TagLogicEnum int32

// List of TagLogicEnum
const (
	TagLogicEnumValue0 TagLogicEnum = 0
	TagLogicEnumValue1 TagLogicEnum = 1
)

// All allowed values of TagLogicEnum enum
var AllowedTagLogicEnumEnumValues = []TagLogicEnum{
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TagLogicEnum) IsValid() bool {
	for _, existing := range AllowedTagLogicEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
