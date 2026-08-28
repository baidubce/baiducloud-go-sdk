package ocr

// VehicleLicenseDirectionEnum the model 'VehicleLicenseDirectionEnum'
type VehicleLicenseDirectionEnum int32

// List of VehicleLicenseDirectionEnum
const (
	VehicleLicenseDirectionEnumValueMinus1 VehicleLicenseDirectionEnum = -1
	VehicleLicenseDirectionEnumValue0      VehicleLicenseDirectionEnum = 0
	VehicleLicenseDirectionEnumValue1      VehicleLicenseDirectionEnum = 1
	VehicleLicenseDirectionEnumValue2      VehicleLicenseDirectionEnum = 2
	VehicleLicenseDirectionEnumValue3      VehicleLicenseDirectionEnum = 3
)

// All allowed values of VehicleLicenseDirectionEnum enum
var AllowedVehicleLicenseDirectionEnumEnumValues = []VehicleLicenseDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VehicleLicenseDirectionEnum) IsValid() bool {
	for _, existing := range AllowedVehicleLicenseDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
