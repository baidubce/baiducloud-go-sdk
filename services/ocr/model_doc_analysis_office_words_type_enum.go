package ocr

// DocAnalysisOfficeWordsTypeEnum the model 'DocAnalysisOfficeWordsTypeEnum'
type DocAnalysisOfficeWordsTypeEnum string

// List of DocAnalysisOfficeWordsTypeEnum
const (
	DocAnalysisOfficeWordsTypeEnumHandwringOnly DocAnalysisOfficeWordsTypeEnum = "handwring_only"
	DocAnalysisOfficeWordsTypeEnumHandprintMix  DocAnalysisOfficeWordsTypeEnum = "handprint_mix"
)

// All allowed values of DocAnalysisOfficeWordsTypeEnum enum
var AllowedDocAnalysisOfficeWordsTypeEnumEnumValues = []DocAnalysisOfficeWordsTypeEnum{
	"handwring_only",
	"handprint_mix",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocAnalysisOfficeWordsTypeEnum) IsValid() bool {
	for _, existing := range AllowedDocAnalysisOfficeWordsTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
