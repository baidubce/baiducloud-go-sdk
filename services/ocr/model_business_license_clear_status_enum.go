package ocr

// BusinessLicenseClearStatusEnum the model 'BusinessLicenseClearStatusEnum'
type BusinessLicenseClearStatusEnum int32

// List of BusinessLicenseClearStatusEnum
const (
	BusinessLicenseClearStatusEnumValue0 BusinessLicenseClearStatusEnum = 0
	BusinessLicenseClearStatusEnumValue1 BusinessLicenseClearStatusEnum = 1
)

// All allowed values of BusinessLicenseClearStatusEnum enum
var AllowedBusinessLicenseClearStatusEnumEnumValues = []BusinessLicenseClearStatusEnum{
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BusinessLicenseClearStatusEnum) IsValid() bool {
	for _, existing := range AllowedBusinessLicenseClearStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
