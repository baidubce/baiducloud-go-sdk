package ocr

// HandwritingLanguageTypeExEnum the model 'HandwritingLanguageTypeExEnum'
type HandwritingLanguageTypeExEnum string

// List of HandwritingLanguageTypeExEnum
const (
	HandwritingLanguageTypeExEnumAutoDetect HandwritingLanguageTypeExEnum = "auto_detect"
	HandwritingLanguageTypeExEnumChnEng     HandwritingLanguageTypeExEnum = "CHN_ENG"
	HandwritingLanguageTypeExEnumEng        HandwritingLanguageTypeExEnum = "ENG"
	HandwritingLanguageTypeExEnumJap        HandwritingLanguageTypeExEnum = "JAP"
	HandwritingLanguageTypeExEnumKor        HandwritingLanguageTypeExEnum = "KOR"
	HandwritingLanguageTypeExEnumFre        HandwritingLanguageTypeExEnum = "FRE"
	HandwritingLanguageTypeExEnumSpa        HandwritingLanguageTypeExEnum = "SPA"
	HandwritingLanguageTypeExEnumPor        HandwritingLanguageTypeExEnum = "POR"
	HandwritingLanguageTypeExEnumGer        HandwritingLanguageTypeExEnum = "GER"
	HandwritingLanguageTypeExEnumIta        HandwritingLanguageTypeExEnum = "ITA"
	HandwritingLanguageTypeExEnumRus        HandwritingLanguageTypeExEnum = "RUS"
	HandwritingLanguageTypeExEnumDan        HandwritingLanguageTypeExEnum = "DAN"
	HandwritingLanguageTypeExEnumDut        HandwritingLanguageTypeExEnum = "DUT"
	HandwritingLanguageTypeExEnumMal        HandwritingLanguageTypeExEnum = "MAL"
	HandwritingLanguageTypeExEnumSwe        HandwritingLanguageTypeExEnum = "SWE"
	HandwritingLanguageTypeExEnumInd        HandwritingLanguageTypeExEnum = "IND"
	HandwritingLanguageTypeExEnumPol        HandwritingLanguageTypeExEnum = "POL"
	HandwritingLanguageTypeExEnumRom        HandwritingLanguageTypeExEnum = "ROM"
	HandwritingLanguageTypeExEnumTur        HandwritingLanguageTypeExEnum = "TUR"
	HandwritingLanguageTypeExEnumGre        HandwritingLanguageTypeExEnum = "GRE"
	HandwritingLanguageTypeExEnumHun        HandwritingLanguageTypeExEnum = "HUN"
	HandwritingLanguageTypeExEnumTha        HandwritingLanguageTypeExEnum = "THA"
	HandwritingLanguageTypeExEnumVie        HandwritingLanguageTypeExEnum = "VIE"
	HandwritingLanguageTypeExEnumAra        HandwritingLanguageTypeExEnum = "ARA"
	HandwritingLanguageTypeExEnumHin        HandwritingLanguageTypeExEnum = "HIN"
)

// All allowed values of HandwritingLanguageTypeExEnum enum
var AllowedHandwritingLanguageTypeExEnumEnumValues = []HandwritingLanguageTypeExEnum{
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
func (v HandwritingLanguageTypeExEnum) IsValid() bool {
	for _, existing := range AllowedHandwritingLanguageTypeExEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
