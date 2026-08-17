package ocr

// MediCheckEnum the model 'MediCheckEnum'
type MediCheckEnum int32

// List of MediCheckEnum
const (
	MediCheckEnumValue0 MediCheckEnum = 0
	MediCheckEnumValue1 MediCheckEnum = 1
)

// All allowed values of MediCheckEnum enum
var AllowedMediCheckEnumEnumValues = []MediCheckEnum{
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MediCheckEnum) IsValid() bool {
	for _, existing := range AllowedMediCheckEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
