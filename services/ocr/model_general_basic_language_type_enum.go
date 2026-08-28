package ocr

// GeneralBasicLanguageTypeEnum the model 'GeneralBasicLanguageTypeEnum'
type GeneralBasicLanguageTypeEnum string

// List of GeneralBasicLanguageTypeEnum
const (
	GeneralBasicLanguageTypeEnumChnEng GeneralBasicLanguageTypeEnum = "CHN_ENG"
	GeneralBasicLanguageTypeEnumEng    GeneralBasicLanguageTypeEnum = "ENG"
	GeneralBasicLanguageTypeEnumJap    GeneralBasicLanguageTypeEnum = "JAP"
	GeneralBasicLanguageTypeEnumKor    GeneralBasicLanguageTypeEnum = "KOR"
	GeneralBasicLanguageTypeEnumFre    GeneralBasicLanguageTypeEnum = "FRE"
	GeneralBasicLanguageTypeEnumSpa    GeneralBasicLanguageTypeEnum = "SPA"
	GeneralBasicLanguageTypeEnumPor    GeneralBasicLanguageTypeEnum = "POR"
	GeneralBasicLanguageTypeEnumGer    GeneralBasicLanguageTypeEnum = "GER"
	GeneralBasicLanguageTypeEnumIta    GeneralBasicLanguageTypeEnum = "ITA"
	GeneralBasicLanguageTypeEnumRus    GeneralBasicLanguageTypeEnum = "RUS"
)

// All allowed values of GeneralBasicLanguageTypeEnum enum
var AllowedGeneralBasicLanguageTypeEnumEnumValues = []GeneralBasicLanguageTypeEnum{
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
func (v GeneralBasicLanguageTypeEnum) IsValid() bool {
	for _, existing := range AllowedGeneralBasicLanguageTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
