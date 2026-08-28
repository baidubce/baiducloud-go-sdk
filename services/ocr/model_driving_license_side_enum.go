package ocr

// DrivingLicenseSideEnum the model 'DrivingLicenseSideEnum'
type DrivingLicenseSideEnum string

// List of DrivingLicenseSideEnum
const (
	DrivingLicenseSideEnumFront DrivingLicenseSideEnum = "front"
	DrivingLicenseSideEnumBack  DrivingLicenseSideEnum = "back"
)

// All allowed values of DrivingLicenseSideEnum enum
var AllowedDrivingLicenseSideEnumEnumValues = []DrivingLicenseSideEnum{
	"front",
	"back",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DrivingLicenseSideEnum) IsValid() bool {
	for _, existing := range AllowedDrivingLicenseSideEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
