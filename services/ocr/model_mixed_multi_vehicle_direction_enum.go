package ocr

// MixedMultiVehicleDirectionEnum the model 'MixedMultiVehicleDirectionEnum'
type MixedMultiVehicleDirectionEnum int32

// List of MixedMultiVehicleDirectionEnum
const (
	MixedMultiVehicleDirectionEnumValueMinus1 MixedMultiVehicleDirectionEnum = -1
	MixedMultiVehicleDirectionEnumValue0      MixedMultiVehicleDirectionEnum = 0
	MixedMultiVehicleDirectionEnumValue1      MixedMultiVehicleDirectionEnum = 1
	MixedMultiVehicleDirectionEnumValue2      MixedMultiVehicleDirectionEnum = 2
	MixedMultiVehicleDirectionEnumValue3      MixedMultiVehicleDirectionEnum = 3
)

// All allowed values of MixedMultiVehicleDirectionEnum enum
var AllowedMixedMultiVehicleDirectionEnumEnumValues = []MixedMultiVehicleDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MixedMultiVehicleDirectionEnum) IsValid() bool {
	for _, existing := range AllowedMixedMultiVehicleDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
