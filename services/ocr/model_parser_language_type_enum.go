package ocr

// ParserLanguageTypeEnum the model 'ParserLanguageTypeEnum'
type ParserLanguageTypeEnum string

// List of ParserLanguageTypeEnum
const (
	ParserLanguageTypeEnumChnEng ParserLanguageTypeEnum = "CHN_ENG"
	ParserLanguageTypeEnumJap    ParserLanguageTypeEnum = "JAP"
	ParserLanguageTypeEnumKor    ParserLanguageTypeEnum = "KOR"
	ParserLanguageTypeEnumFre    ParserLanguageTypeEnum = "FRE"
	ParserLanguageTypeEnumSpa    ParserLanguageTypeEnum = "SPA"
	ParserLanguageTypeEnumPor    ParserLanguageTypeEnum = "POR"
	ParserLanguageTypeEnumGer    ParserLanguageTypeEnum = "GER"
	ParserLanguageTypeEnumIta    ParserLanguageTypeEnum = "ITA"
	ParserLanguageTypeEnumRus    ParserLanguageTypeEnum = "RUS"
	ParserLanguageTypeEnumDan    ParserLanguageTypeEnum = "DAN"
	ParserLanguageTypeEnumDut    ParserLanguageTypeEnum = "DUT"
	ParserLanguageTypeEnumMal    ParserLanguageTypeEnum = "MAL"
	ParserLanguageTypeEnumSwe    ParserLanguageTypeEnum = "SWE"
	ParserLanguageTypeEnumInd    ParserLanguageTypeEnum = "IND"
	ParserLanguageTypeEnumPol    ParserLanguageTypeEnum = "POL"
	ParserLanguageTypeEnumRom    ParserLanguageTypeEnum = "ROM"
	ParserLanguageTypeEnumTur    ParserLanguageTypeEnum = "TUR"
	ParserLanguageTypeEnumGre    ParserLanguageTypeEnum = "GRE"
	ParserLanguageTypeEnumHun    ParserLanguageTypeEnum = "HUN"
	ParserLanguageTypeEnumTha    ParserLanguageTypeEnum = "THA"
	ParserLanguageTypeEnumVie    ParserLanguageTypeEnum = "VIE"
	ParserLanguageTypeEnumAra    ParserLanguageTypeEnum = "ARA"
	ParserLanguageTypeEnumHin    ParserLanguageTypeEnum = "HIN"
)

// All allowed values of ParserLanguageTypeEnum enum
var AllowedParserLanguageTypeEnumEnumValues = []ParserLanguageTypeEnum{
	"CHN_ENG",
	"JAP",
	"KOR",
	"FRE",
	"SPA",
	"POR",
	"GER",
	"ITA",
	"RUS",
	"DAN",
	"DUT",
	"MAL",
	"SWE",
	"IND",
	"POL",
	"ROM",
	"TUR",
	"GRE",
	"HUN",
	"THA",
	"VIE",
	"ARA",
	"HIN",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ParserLanguageTypeEnum) IsValid() bool {
	for _, existing := range AllowedParserLanguageTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
