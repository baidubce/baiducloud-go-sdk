package ocr

// BusinessLicenseRiskTypeEnum the model 'BusinessLicenseRiskTypeEnum'
type BusinessLicenseRiskTypeEnum string

// List of BusinessLicenseRiskTypeEnum
const (
	BusinessLicenseRiskTypeEnumNormal BusinessLicenseRiskTypeEnum = "normal"
	BusinessLicenseRiskTypeEnumCopy   BusinessLicenseRiskTypeEnum = "copy"
	BusinessLicenseRiskTypeEnumScreen BusinessLicenseRiskTypeEnum = "screen"
	BusinessLicenseRiskTypeEnumScan   BusinessLicenseRiskTypeEnum = "scan"
)

// All allowed values of BusinessLicenseRiskTypeEnum enum
var AllowedBusinessLicenseRiskTypeEnumEnumValues = []BusinessLicenseRiskTypeEnum{
	"normal",
	"copy",
	"screen",
	"scan",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BusinessLicenseRiskTypeEnum) IsValid() bool {
	for _, existing := range AllowedBusinessLicenseRiskTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
