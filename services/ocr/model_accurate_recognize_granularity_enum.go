package ocr

// AccurateRecognizeGranularityEnum the model 'AccurateRecognizeGranularityEnum'
type AccurateRecognizeGranularityEnum string

// List of AccurateRecognizeGranularityEnum
const (
	AccurateRecognizeGranularityEnumBig   AccurateRecognizeGranularityEnum = "big"
	AccurateRecognizeGranularityEnumSmall AccurateRecognizeGranularityEnum = "small"
)

// All allowed values of AccurateRecognizeGranularityEnum enum
var AllowedAccurateRecognizeGranularityEnumEnumValues = []AccurateRecognizeGranularityEnum{
	"big",
	"small",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccurateRecognizeGranularityEnum) IsValid() bool {
	for _, existing := range AllowedAccurateRecognizeGranularityEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
