package ocr

// EducationPaperCutEduWordsTypeEnum the model 'EducationPaperCutEduWordsTypeEnum'
type EducationPaperCutEduWordsTypeEnum string

// List of EducationPaperCutEduWordsTypeEnum
const (
	EducationPaperCutEduWordsTypeEnumHandwringOnly EducationPaperCutEduWordsTypeEnum = "handwring_only"
	EducationPaperCutEduWordsTypeEnumHandprintMix  EducationPaperCutEduWordsTypeEnum = "handprint_mix"
)

// All allowed values of EducationPaperCutEduWordsTypeEnum enum
var AllowedEducationPaperCutEduWordsTypeEnumEnumValues = []EducationPaperCutEduWordsTypeEnum{
	"handwring_only",
	"handprint_mix",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EducationPaperCutEduWordsTypeEnum) IsValid() bool {
	for _, existing := range AllowedEducationPaperCutEduWordsTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
