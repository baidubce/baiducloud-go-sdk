package ocr

// PaperCutLanguageTypeEnum the model 'PaperCutLanguageTypeEnum'
type PaperCutLanguageTypeEnum string

// List of PaperCutLanguageTypeEnum
const (
	PaperCutLanguageTypeEnumChnEng PaperCutLanguageTypeEnum = "CHN_ENG"
	PaperCutLanguageTypeEnumEng    PaperCutLanguageTypeEnum = "ENG"
)

// All allowed values of PaperCutLanguageTypeEnum enum
var AllowedPaperCutLanguageTypeEnumEnumValues = []PaperCutLanguageTypeEnum{
	"CHN_ENG",
	"ENG",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaperCutLanguageTypeEnum) IsValid() bool {
	for _, existing := range AllowedPaperCutLanguageTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
