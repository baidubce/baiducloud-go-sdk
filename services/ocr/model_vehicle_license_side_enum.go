package ocr

// VehicleLicenseSideEnum the model 'VehicleLicenseSideEnum'
type VehicleLicenseSideEnum string

// List of VehicleLicenseSideEnum
const (
	VehicleLicenseSideEnumFront VehicleLicenseSideEnum = "front"
	VehicleLicenseSideEnumBack  VehicleLicenseSideEnum = "back"
)

// All allowed values of VehicleLicenseSideEnum enum
var AllowedVehicleLicenseSideEnumEnumValues = []VehicleLicenseSideEnum{
	"front",
	"back",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VehicleLicenseSideEnum) IsValid() bool {
	for _, existing := range AllowedVehicleLicenseSideEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
