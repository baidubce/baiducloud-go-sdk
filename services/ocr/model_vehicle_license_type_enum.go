package ocr

// VehicleLicenseTypeEnum the model 'VehicleLicenseTypeEnum'
type VehicleLicenseTypeEnum string

// List of VehicleLicenseTypeEnum
const (
	VehicleLicenseTypeEnumVehicleFront VehicleLicenseTypeEnum = "vehicle_front"
	VehicleLicenseTypeEnumVehicleBack  VehicleLicenseTypeEnum = "vehicle_back"
	VehicleLicenseTypeEnumDrivingFront VehicleLicenseTypeEnum = "driving_front"
	VehicleLicenseTypeEnumDrivingBack  VehicleLicenseTypeEnum = "driving_back"
)

// All allowed values of VehicleLicenseTypeEnum enum
var AllowedVehicleLicenseTypeEnumEnumValues = []VehicleLicenseTypeEnum{
	"vehicle_front",
	"vehicle_back",
	"driving_front",
	"driving_back",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VehicleLicenseTypeEnum) IsValid() bool {
	for _, existing := range AllowedVehicleLicenseTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
