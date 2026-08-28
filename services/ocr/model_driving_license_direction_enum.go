package ocr

// DrivingLicenseDirectionEnum the model 'DrivingLicenseDirectionEnum'
type DrivingLicenseDirectionEnum int32

// List of DrivingLicenseDirectionEnum
const (
	DrivingLicenseDirectionEnumValueMinus1 DrivingLicenseDirectionEnum = -1
	DrivingLicenseDirectionEnumValue0      DrivingLicenseDirectionEnum = 0
	DrivingLicenseDirectionEnumValue1      DrivingLicenseDirectionEnum = 1
	DrivingLicenseDirectionEnumValue2      DrivingLicenseDirectionEnum = 2
	DrivingLicenseDirectionEnumValue3      DrivingLicenseDirectionEnum = 3
)

// All allowed values of DrivingLicenseDirectionEnum enum
var AllowedDrivingLicenseDirectionEnumEnumValues = []DrivingLicenseDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DrivingLicenseDirectionEnum) IsValid() bool {
	for _, existing := range AllowedDrivingLicenseDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
