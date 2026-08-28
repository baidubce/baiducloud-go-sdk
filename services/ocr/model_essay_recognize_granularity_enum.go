package ocr

// EssayRecognizeGranularityEnum the model 'EssayRecognizeGranularityEnum'
type EssayRecognizeGranularityEnum string

// List of EssayRecognizeGranularityEnum
const (
	EssayRecognizeGranularityEnumLine EssayRecognizeGranularityEnum = "line"
	EssayRecognizeGranularityEnumWord EssayRecognizeGranularityEnum = "word"
	EssayRecognizeGranularityEnumNone EssayRecognizeGranularityEnum = "none"
)

// All allowed values of EssayRecognizeGranularityEnum enum
var AllowedEssayRecognizeGranularityEnumEnumValues = []EssayRecognizeGranularityEnum{
	"line",
	"word",
	"none",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EssayRecognizeGranularityEnum) IsValid() bool {
	for _, existing := range AllowedEssayRecognizeGranularityEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
