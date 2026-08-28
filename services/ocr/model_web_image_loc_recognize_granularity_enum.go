package ocr

// WebImageLocRecognizeGranularityEnum the model 'WebImageLocRecognizeGranularityEnum'
type WebImageLocRecognizeGranularityEnum string

// List of WebImageLocRecognizeGranularityEnum
const (
	WebImageLocRecognizeGranularityEnumBig   WebImageLocRecognizeGranularityEnum = "big"
	WebImageLocRecognizeGranularityEnumSmall WebImageLocRecognizeGranularityEnum = "small"
)

// All allowed values of WebImageLocRecognizeGranularityEnum enum
var AllowedWebImageLocRecognizeGranularityEnumEnumValues = []WebImageLocRecognizeGranularityEnum{
	"big",
	"small",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebImageLocRecognizeGranularityEnum) IsValid() bool {
	for _, existing := range AllowedWebImageLocRecognizeGranularityEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
