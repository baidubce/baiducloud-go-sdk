package ocr

// MultiIdcardRiskTypeEnum the model 'MultiIdcardRiskTypeEnum'
type MultiIdcardRiskTypeEnum string

// List of MultiIdcardRiskTypeEnum
const (
	MultiIdcardRiskTypeEnumNormal MultiIdcardRiskTypeEnum = "normal"
	MultiIdcardRiskTypeEnumCopy   MultiIdcardRiskTypeEnum = "copy"
	MultiIdcardRiskTypeEnumScreen MultiIdcardRiskTypeEnum = "screen"
	MultiIdcardRiskTypeEnumScan   MultiIdcardRiskTypeEnum = "scan"
)

// All allowed values of MultiIdcardRiskTypeEnum enum
var AllowedMultiIdcardRiskTypeEnumEnumValues = []MultiIdcardRiskTypeEnum{
	"normal",
	"copy",
	"screen",
	"scan",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MultiIdcardRiskTypeEnum) IsValid() bool {
	for _, existing := range AllowedMultiIdcardRiskTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
