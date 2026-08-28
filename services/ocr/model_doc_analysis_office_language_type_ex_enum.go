package ocr

// DocAnalysisOfficeLanguageTypeExEnum the model 'DocAnalysisOfficeLanguageTypeExEnum'
type DocAnalysisOfficeLanguageTypeExEnum string

// List of DocAnalysisOfficeLanguageTypeExEnum
const (
	DocAnalysisOfficeLanguageTypeExEnumAutoDetect DocAnalysisOfficeLanguageTypeExEnum = "auto_detect"
	DocAnalysisOfficeLanguageTypeExEnumChnEng     DocAnalysisOfficeLanguageTypeExEnum = "CHN_ENG"
	DocAnalysisOfficeLanguageTypeExEnumEng        DocAnalysisOfficeLanguageTypeExEnum = "ENG"
	DocAnalysisOfficeLanguageTypeExEnumJap        DocAnalysisOfficeLanguageTypeExEnum = "JAP"
	DocAnalysisOfficeLanguageTypeExEnumKor        DocAnalysisOfficeLanguageTypeExEnum = "KOR"
	DocAnalysisOfficeLanguageTypeExEnumFre        DocAnalysisOfficeLanguageTypeExEnum = "FRE"
	DocAnalysisOfficeLanguageTypeExEnumSpa        DocAnalysisOfficeLanguageTypeExEnum = "SPA"
	DocAnalysisOfficeLanguageTypeExEnumPor        DocAnalysisOfficeLanguageTypeExEnum = "POR"
	DocAnalysisOfficeLanguageTypeExEnumGer        DocAnalysisOfficeLanguageTypeExEnum = "GER"
	DocAnalysisOfficeLanguageTypeExEnumIta        DocAnalysisOfficeLanguageTypeExEnum = "ITA"
	DocAnalysisOfficeLanguageTypeExEnumRus        DocAnalysisOfficeLanguageTypeExEnum = "RUS"
	DocAnalysisOfficeLanguageTypeExEnumDan        DocAnalysisOfficeLanguageTypeExEnum = "DAN"
	DocAnalysisOfficeLanguageTypeExEnumDut        DocAnalysisOfficeLanguageTypeExEnum = "DUT"
	DocAnalysisOfficeLanguageTypeExEnumMal        DocAnalysisOfficeLanguageTypeExEnum = "MAL"
	DocAnalysisOfficeLanguageTypeExEnumSwe        DocAnalysisOfficeLanguageTypeExEnum = "SWE"
	DocAnalysisOfficeLanguageTypeExEnumInd        DocAnalysisOfficeLanguageTypeExEnum = "IND"
	DocAnalysisOfficeLanguageTypeExEnumPol        DocAnalysisOfficeLanguageTypeExEnum = "POL"
	DocAnalysisOfficeLanguageTypeExEnumRom        DocAnalysisOfficeLanguageTypeExEnum = "ROM"
	DocAnalysisOfficeLanguageTypeExEnumTur        DocAnalysisOfficeLanguageTypeExEnum = "TUR"
	DocAnalysisOfficeLanguageTypeExEnumGre        DocAnalysisOfficeLanguageTypeExEnum = "GRE"
	DocAnalysisOfficeLanguageTypeExEnumHun        DocAnalysisOfficeLanguageTypeExEnum = "HUN"
	DocAnalysisOfficeLanguageTypeExEnumTha        DocAnalysisOfficeLanguageTypeExEnum = "THA"
	DocAnalysisOfficeLanguageTypeExEnumVie        DocAnalysisOfficeLanguageTypeExEnum = "VIE"
	DocAnalysisOfficeLanguageTypeExEnumAra        DocAnalysisOfficeLanguageTypeExEnum = "ARA"
	DocAnalysisOfficeLanguageTypeExEnumHin        DocAnalysisOfficeLanguageTypeExEnum = "HIN"
)

// All allowed values of DocAnalysisOfficeLanguageTypeExEnum enum
var AllowedDocAnalysisOfficeLanguageTypeExEnumEnumValues = []DocAnalysisOfficeLanguageTypeExEnum{
	"auto_detect",
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
func (v DocAnalysisOfficeLanguageTypeExEnum) IsValid() bool {
	for _, existing := range AllowedDocAnalysisOfficeLanguageTypeExEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
