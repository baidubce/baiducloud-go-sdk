package ocr

// NumbersRecognizeGranularityEnum the model 'NumbersRecognizeGranularityEnum'
type NumbersRecognizeGranularityEnum string

// List of NumbersRecognizeGranularityEnum
const (
	NumbersRecognizeGranularityEnumBig   NumbersRecognizeGranularityEnum = "big"
	NumbersRecognizeGranularityEnumSmall NumbersRecognizeGranularityEnum = "small"
)

// All allowed values of NumbersRecognizeGranularityEnum enum
var AllowedNumbersRecognizeGranularityEnumEnumValues = []NumbersRecognizeGranularityEnum{
	"big",
	"small",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NumbersRecognizeGranularityEnum) IsValid() bool {
	for _, existing := range AllowedNumbersRecognizeGranularityEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
