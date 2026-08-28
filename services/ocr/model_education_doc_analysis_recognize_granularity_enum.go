package ocr

// EducationDocAnalysisRecognizeGranularityEnum the model 'EducationDocAnalysisRecognizeGranularityEnum'
type EducationDocAnalysisRecognizeGranularityEnum string

// List of EducationDocAnalysisRecognizeGranularityEnum
const (
	EducationDocAnalysisRecognizeGranularityEnumBig   EducationDocAnalysisRecognizeGranularityEnum = "big"
	EducationDocAnalysisRecognizeGranularityEnumSmall EducationDocAnalysisRecognizeGranularityEnum = "small"
)

// All allowed values of EducationDocAnalysisRecognizeGranularityEnum enum
var AllowedEducationDocAnalysisRecognizeGranularityEnumEnumValues = []EducationDocAnalysisRecognizeGranularityEnum{
	"big",
	"small",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EducationDocAnalysisRecognizeGranularityEnum) IsValid() bool {
	for _, existing := range AllowedEducationDocAnalysisRecognizeGranularityEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
