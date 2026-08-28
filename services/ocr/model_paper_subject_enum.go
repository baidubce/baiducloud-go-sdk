package ocr

// PaperSubjectEnum the model 'PaperSubjectEnum'
type PaperSubjectEnum string

// List of PaperSubjectEnum
const (
	PaperSubjectEnumChinese   PaperSubjectEnum = "chinese"
	PaperSubjectEnumMath      PaperSubjectEnum = "math"
	PaperSubjectEnumEnglish   PaperSubjectEnum = "english"
	PaperSubjectEnumPhysics   PaperSubjectEnum = "physics"
	PaperSubjectEnumChemistry PaperSubjectEnum = "chemistry"
	PaperSubjectEnumBiology   PaperSubjectEnum = "biology"
	PaperSubjectEnumHistory   PaperSubjectEnum = "history"
	PaperSubjectEnumGeography PaperSubjectEnum = "geography"
	PaperSubjectEnumPolitics  PaperSubjectEnum = "politics"
)

// All allowed values of PaperSubjectEnum enum
var AllowedPaperSubjectEnumEnumValues = []PaperSubjectEnum{
	"chinese",
	"math",
	"english",
	"physics",
	"chemistry",
	"biology",
	"history",
	"geography",
	"politics",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaperSubjectEnum) IsValid() bool {
	for _, existing := range AllowedPaperSubjectEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
