package ocr

// GeneralRecognizeGranularityEnum the model 'GeneralRecognizeGranularityEnum'
type GeneralRecognizeGranularityEnum string

// List of GeneralRecognizeGranularityEnum
const (
	GeneralRecognizeGranularityEnumBig   GeneralRecognizeGranularityEnum = "big"
	GeneralRecognizeGranularityEnumSmall GeneralRecognizeGranularityEnum = "small"
)

// All allowed values of GeneralRecognizeGranularityEnum enum
var AllowedGeneralRecognizeGranularityEnumEnumValues = []GeneralRecognizeGranularityEnum{
	"big",
	"small",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GeneralRecognizeGranularityEnum) IsValid() bool {
	for _, existing := range AllowedGeneralRecognizeGranularityEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
