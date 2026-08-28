package ocr

// VehicleRegistrationCertificateStringDirectionEnum the model 'VehicleRegistrationCertificateStringDirectionEnum'
type VehicleRegistrationCertificateStringDirectionEnum string

// List of VehicleRegistrationCertificateStringDirectionEnum
const (
	VehicleRegistrationCertificateStringDirectionEnumValueMinus1 VehicleRegistrationCertificateStringDirectionEnum = "-1"
	VehicleRegistrationCertificateStringDirectionEnumValue0      VehicleRegistrationCertificateStringDirectionEnum = "0"
	VehicleRegistrationCertificateStringDirectionEnumValue1      VehicleRegistrationCertificateStringDirectionEnum = "1"
	VehicleRegistrationCertificateStringDirectionEnumValue2      VehicleRegistrationCertificateStringDirectionEnum = "2"
	VehicleRegistrationCertificateStringDirectionEnumValue3      VehicleRegistrationCertificateStringDirectionEnum = "3"
)

// All allowed values of VehicleRegistrationCertificateStringDirectionEnum enum
var AllowedVehicleRegistrationCertificateStringDirectionEnumEnumValues = []VehicleRegistrationCertificateStringDirectionEnum{
	"-1",
	"0",
	"1",
	"2",
	"3",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VehicleRegistrationCertificateStringDirectionEnum) IsValid() bool {
	for _, existing := range AllowedVehicleRegistrationCertificateStringDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
