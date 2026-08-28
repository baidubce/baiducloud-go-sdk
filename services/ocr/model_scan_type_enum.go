package ocr

// ScanTypeEnum the model 'ScanTypeEnum'
type ScanTypeEnum int32

// List of ScanTypeEnum
const (
	ScanTypeEnumValue1 ScanTypeEnum = 1
	ScanTypeEnumValue2 ScanTypeEnum = 2
	ScanTypeEnumValue3 ScanTypeEnum = 3
)

// All allowed values of ScanTypeEnum enum
var AllowedScanTypeEnumEnumValues = []ScanTypeEnum{
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScanTypeEnum) IsValid() bool {
	for _, existing := range AllowedScanTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
