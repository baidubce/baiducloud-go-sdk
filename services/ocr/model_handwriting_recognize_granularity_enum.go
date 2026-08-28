package ocr

// HandwritingRecognizeGranularityEnum the model 'HandwritingRecognizeGranularityEnum'
type HandwritingRecognizeGranularityEnum string

// List of HandwritingRecognizeGranularityEnum
const (
	HandwritingRecognizeGranularityEnumBig   HandwritingRecognizeGranularityEnum = "big"
	HandwritingRecognizeGranularityEnumSmall HandwritingRecognizeGranularityEnum = "small"
)

// All allowed values of HandwritingRecognizeGranularityEnum enum
var AllowedHandwritingRecognizeGranularityEnumEnumValues = []HandwritingRecognizeGranularityEnum{
	"big",
	"small",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HandwritingRecognizeGranularityEnum) IsValid() bool {
	for _, existing := range AllowedHandwritingRecognizeGranularityEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
