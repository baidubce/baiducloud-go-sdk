package ocr

// DocAnalysisOfficeRecognizeGranularityEnum the model 'DocAnalysisOfficeRecognizeGranularityEnum'
type DocAnalysisOfficeRecognizeGranularityEnum string

// List of DocAnalysisOfficeRecognizeGranularityEnum
const (
	DocAnalysisOfficeRecognizeGranularityEnumBig   DocAnalysisOfficeRecognizeGranularityEnum = "big"
	DocAnalysisOfficeRecognizeGranularityEnumSmall DocAnalysisOfficeRecognizeGranularityEnum = "small"
)

// All allowed values of DocAnalysisOfficeRecognizeGranularityEnum enum
var AllowedDocAnalysisOfficeRecognizeGranularityEnumEnumValues = []DocAnalysisOfficeRecognizeGranularityEnum{
	"big",
	"small",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocAnalysisOfficeRecognizeGranularityEnum) IsValid() bool {
	for _, existing := range AllowedDocAnalysisOfficeRecognizeGranularityEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
