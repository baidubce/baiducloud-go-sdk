package ocr

// GeneralLanguageTypeEnum the model 'GeneralLanguageTypeEnum'
type GeneralLanguageTypeEnum string

// List of GeneralLanguageTypeEnum
const (
	GeneralLanguageTypeEnumChnEng GeneralLanguageTypeEnum = "CHN_ENG"
	GeneralLanguageTypeEnumEng    GeneralLanguageTypeEnum = "ENG"
	GeneralLanguageTypeEnumJap    GeneralLanguageTypeEnum = "JAP"
	GeneralLanguageTypeEnumKor    GeneralLanguageTypeEnum = "KOR"
	GeneralLanguageTypeEnumFre    GeneralLanguageTypeEnum = "FRE"
	GeneralLanguageTypeEnumSpa    GeneralLanguageTypeEnum = "SPA"
	GeneralLanguageTypeEnumPor    GeneralLanguageTypeEnum = "POR"
	GeneralLanguageTypeEnumGer    GeneralLanguageTypeEnum = "GER"
	GeneralLanguageTypeEnumIta    GeneralLanguageTypeEnum = "ITA"
	GeneralLanguageTypeEnumRus    GeneralLanguageTypeEnum = "RUS"
)

// All allowed values of GeneralLanguageTypeEnum enum
var AllowedGeneralLanguageTypeEnumEnumValues = []GeneralLanguageTypeEnum{
	"CHN_ENG",
	"ENG",
	"JAP",
	"KOR",
	"FRE",
	"SPA",
	"POR",
	"GER",
	"ITA",
	"RUS",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GeneralLanguageTypeEnum) IsValid() bool {
	for _, existing := range AllowedGeneralLanguageTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
