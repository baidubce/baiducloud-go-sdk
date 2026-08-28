package ocr

// BusinessLicenseCompleteStatusEnum the model 'BusinessLicenseCompleteStatusEnum'
type BusinessLicenseCompleteStatusEnum int32

// List of BusinessLicenseCompleteStatusEnum
const (
	BusinessLicenseCompleteStatusEnumValue0 BusinessLicenseCompleteStatusEnum = 0
	BusinessLicenseCompleteStatusEnumValue1 BusinessLicenseCompleteStatusEnum = 1
)

// All allowed values of BusinessLicenseCompleteStatusEnum enum
var AllowedBusinessLicenseCompleteStatusEnumEnumValues = []BusinessLicenseCompleteStatusEnum{
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BusinessLicenseCompleteStatusEnum) IsValid() bool {
	for _, existing := range AllowedBusinessLicenseCompleteStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
