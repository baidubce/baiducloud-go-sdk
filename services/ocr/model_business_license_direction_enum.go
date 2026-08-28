package ocr

// BusinessLicenseDirectionEnum the model 'BusinessLicenseDirectionEnum'
type BusinessLicenseDirectionEnum int32

// List of BusinessLicenseDirectionEnum
const (
	BusinessLicenseDirectionEnumValueMinus1 BusinessLicenseDirectionEnum = -1
	BusinessLicenseDirectionEnumValue0      BusinessLicenseDirectionEnum = 0
	BusinessLicenseDirectionEnumValue1      BusinessLicenseDirectionEnum = 1
	BusinessLicenseDirectionEnumValue2      BusinessLicenseDirectionEnum = 2
	BusinessLicenseDirectionEnumValue3      BusinessLicenseDirectionEnum = 3
)

// All allowed values of BusinessLicenseDirectionEnum enum
var AllowedBusinessLicenseDirectionEnumEnumValues = []BusinessLicenseDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BusinessLicenseDirectionEnum) IsValid() bool {
	for _, existing := range AllowedBusinessLicenseDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
