package ocr

// AccurateBasicLanguageTypeExEnum the model 'AccurateBasicLanguageTypeExEnum'
type AccurateBasicLanguageTypeExEnum string

// List of AccurateBasicLanguageTypeExEnum
const (
	AccurateBasicLanguageTypeExEnumAutoDetect AccurateBasicLanguageTypeExEnum = "auto_detect"
	AccurateBasicLanguageTypeExEnumChnEng     AccurateBasicLanguageTypeExEnum = "CHN_ENG"
	AccurateBasicLanguageTypeExEnumEng        AccurateBasicLanguageTypeExEnum = "ENG"
	AccurateBasicLanguageTypeExEnumJap        AccurateBasicLanguageTypeExEnum = "JAP"
	AccurateBasicLanguageTypeExEnumKor        AccurateBasicLanguageTypeExEnum = "KOR"
	AccurateBasicLanguageTypeExEnumFre        AccurateBasicLanguageTypeExEnum = "FRE"
	AccurateBasicLanguageTypeExEnumSpa        AccurateBasicLanguageTypeExEnum = "SPA"
	AccurateBasicLanguageTypeExEnumPor        AccurateBasicLanguageTypeExEnum = "POR"
	AccurateBasicLanguageTypeExEnumGer        AccurateBasicLanguageTypeExEnum = "GER"
	AccurateBasicLanguageTypeExEnumIta        AccurateBasicLanguageTypeExEnum = "ITA"
	AccurateBasicLanguageTypeExEnumRus        AccurateBasicLanguageTypeExEnum = "RUS"
	AccurateBasicLanguageTypeExEnumDan        AccurateBasicLanguageTypeExEnum = "DAN"
	AccurateBasicLanguageTypeExEnumDut        AccurateBasicLanguageTypeExEnum = "DUT"
	AccurateBasicLanguageTypeExEnumMal        AccurateBasicLanguageTypeExEnum = "MAL"
	AccurateBasicLanguageTypeExEnumSwe        AccurateBasicLanguageTypeExEnum = "SWE"
	AccurateBasicLanguageTypeExEnumInd        AccurateBasicLanguageTypeExEnum = "IND"
	AccurateBasicLanguageTypeExEnumPol        AccurateBasicLanguageTypeExEnum = "POL"
	AccurateBasicLanguageTypeExEnumRom        AccurateBasicLanguageTypeExEnum = "ROM"
	AccurateBasicLanguageTypeExEnumTur        AccurateBasicLanguageTypeExEnum = "TUR"
	AccurateBasicLanguageTypeExEnumGre        AccurateBasicLanguageTypeExEnum = "GRE"
	AccurateBasicLanguageTypeExEnumHun        AccurateBasicLanguageTypeExEnum = "HUN"
	AccurateBasicLanguageTypeExEnumTha        AccurateBasicLanguageTypeExEnum = "THA"
	AccurateBasicLanguageTypeExEnumVie        AccurateBasicLanguageTypeExEnum = "VIE"
	AccurateBasicLanguageTypeExEnumAra        AccurateBasicLanguageTypeExEnum = "ARA"
	AccurateBasicLanguageTypeExEnumHin        AccurateBasicLanguageTypeExEnum = "HIN"
)

// All allowed values of AccurateBasicLanguageTypeExEnum enum
var AllowedAccurateBasicLanguageTypeExEnumEnumValues = []AccurateBasicLanguageTypeExEnum{
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
func (v AccurateBasicLanguageTypeExEnum) IsValid() bool {
	for _, existing := range AllowedAccurateBasicLanguageTypeExEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
