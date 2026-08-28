package ocr

// VehicleCertificateDirectionEnum the model 'VehicleCertificateDirectionEnum'
type VehicleCertificateDirectionEnum int32

// List of VehicleCertificateDirectionEnum
const (
	VehicleCertificateDirectionEnumValueMinus1 VehicleCertificateDirectionEnum = -1
	VehicleCertificateDirectionEnumValue0      VehicleCertificateDirectionEnum = 0
	VehicleCertificateDirectionEnumValue1      VehicleCertificateDirectionEnum = 1
	VehicleCertificateDirectionEnumValue2      VehicleCertificateDirectionEnum = 2
	VehicleCertificateDirectionEnumValue3      VehicleCertificateDirectionEnum = 3
)

// All allowed values of VehicleCertificateDirectionEnum enum
var AllowedVehicleCertificateDirectionEnumEnumValues = []VehicleCertificateDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VehicleCertificateDirectionEnum) IsValid() bool {
	for _, existing := range AllowedVehicleCertificateDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
