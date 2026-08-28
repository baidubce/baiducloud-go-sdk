package ocr

// EducationDocAnalysisWordsTypeEnum the model 'EducationDocAnalysisWordsTypeEnum'
type EducationDocAnalysisWordsTypeEnum string

// List of EducationDocAnalysisWordsTypeEnum
const (
	EducationDocAnalysisWordsTypeEnumHandwringOnly EducationDocAnalysisWordsTypeEnum = "handwring_only"
	EducationDocAnalysisWordsTypeEnumHandprintMix  EducationDocAnalysisWordsTypeEnum = "handprint_mix"
)

// All allowed values of EducationDocAnalysisWordsTypeEnum enum
var AllowedEducationDocAnalysisWordsTypeEnumEnumValues = []EducationDocAnalysisWordsTypeEnum{
	"handwring_only",
	"handprint_mix",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EducationDocAnalysisWordsTypeEnum) IsValid() bool {
	for _, existing := range AllowedEducationDocAnalysisWordsTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
