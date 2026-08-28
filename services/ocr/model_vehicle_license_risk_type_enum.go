package ocr

// VehicleLicenseRiskTypeEnum the model 'VehicleLicenseRiskTypeEnum'
type VehicleLicenseRiskTypeEnum string

// List of VehicleLicenseRiskTypeEnum
const (
	VehicleLicenseRiskTypeEnumNormal VehicleLicenseRiskTypeEnum = "normal"
	VehicleLicenseRiskTypeEnumCopy   VehicleLicenseRiskTypeEnum = "copy"
	VehicleLicenseRiskTypeEnumScreen VehicleLicenseRiskTypeEnum = "screen"
)

// All allowed values of VehicleLicenseRiskTypeEnum enum
var AllowedVehicleLicenseRiskTypeEnumEnumValues = []VehicleLicenseRiskTypeEnum{
	"normal",
	"copy",
	"screen",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VehicleLicenseRiskTypeEnum) IsValid() bool {
	for _, existing := range AllowedVehicleLicenseRiskTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
