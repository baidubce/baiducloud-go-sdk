package ocr

// AccurateLanguageTypeExEnum the model 'AccurateLanguageTypeExEnum'
type AccurateLanguageTypeExEnum string

// List of AccurateLanguageTypeExEnum
const (
	AccurateLanguageTypeExEnumAutoDetect AccurateLanguageTypeExEnum = "auto_detect"
	AccurateLanguageTypeExEnumChnEng     AccurateLanguageTypeExEnum = "CHN_ENG"
	AccurateLanguageTypeExEnumEng        AccurateLanguageTypeExEnum = "ENG"
	AccurateLanguageTypeExEnumJap        AccurateLanguageTypeExEnum = "JAP"
	AccurateLanguageTypeExEnumKor        AccurateLanguageTypeExEnum = "KOR"
	AccurateLanguageTypeExEnumFre        AccurateLanguageTypeExEnum = "FRE"
	AccurateLanguageTypeExEnumSpa        AccurateLanguageTypeExEnum = "SPA"
	AccurateLanguageTypeExEnumPor        AccurateLanguageTypeExEnum = "POR"
	AccurateLanguageTypeExEnumGer        AccurateLanguageTypeExEnum = "GER"
	AccurateLanguageTypeExEnumIta        AccurateLanguageTypeExEnum = "ITA"
	AccurateLanguageTypeExEnumRus        AccurateLanguageTypeExEnum = "RUS"
	AccurateLanguageTypeExEnumDan        AccurateLanguageTypeExEnum = "DAN"
	AccurateLanguageTypeExEnumDut        AccurateLanguageTypeExEnum = "DUT"
	AccurateLanguageTypeExEnumMal        AccurateLanguageTypeExEnum = "MAL"
	AccurateLanguageTypeExEnumSwe        AccurateLanguageTypeExEnum = "SWE"
	AccurateLanguageTypeExEnumInd        AccurateLanguageTypeExEnum = "IND"
	AccurateLanguageTypeExEnumPol        AccurateLanguageTypeExEnum = "POL"
	AccurateLanguageTypeExEnumRom        AccurateLanguageTypeExEnum = "ROM"
	AccurateLanguageTypeExEnumTur        AccurateLanguageTypeExEnum = "TUR"
	AccurateLanguageTypeExEnumGre        AccurateLanguageTypeExEnum = "GRE"
	AccurateLanguageTypeExEnumHun        AccurateLanguageTypeExEnum = "HUN"
	AccurateLanguageTypeExEnumTha        AccurateLanguageTypeExEnum = "THA"
	AccurateLanguageTypeExEnumVie        AccurateLanguageTypeExEnum = "VIE"
	AccurateLanguageTypeExEnumAra        AccurateLanguageTypeExEnum = "ARA"
	AccurateLanguageTypeExEnumHin        AccurateLanguageTypeExEnum = "HIN"
)

// All allowed values of AccurateLanguageTypeExEnum enum
var AllowedAccurateLanguageTypeExEnumEnumValues = []AccurateLanguageTypeExEnum{
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
func (v AccurateLanguageTypeExEnum) IsValid() bool {
	for _, existing := range AllowedAccurateLanguageTypeExEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
