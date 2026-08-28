package ocr

// AccuracyEnum the model 'AccuracyEnum'
type AccuracyEnum string

// List of AccuracyEnum
const (
	AccuracyEnumNormal AccuracyEnum = "normal"
	AccuracyEnumHigh   AccuracyEnum = "high"
)

// All allowed values of AccuracyEnum enum
var AllowedAccuracyEnumEnumValues = []AccuracyEnum{
	"normal",
	"high",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccuracyEnum) IsValid() bool {
	for _, existing := range AllowedAccuracyEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
