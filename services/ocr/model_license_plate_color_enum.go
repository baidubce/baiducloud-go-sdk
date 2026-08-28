package ocr

// LicensePlateColorEnum the model 'LicensePlateColorEnum'
type LicensePlateColorEnum string

// List of LicensePlateColorEnum
const (
	LicensePlateColorEnumBlue        LicensePlateColorEnum = "blue"
	LicensePlateColorEnumGreen       LicensePlateColorEnum = "green"
	LicensePlateColorEnumYellow      LicensePlateColorEnum = "yellow"
	LicensePlateColorEnumWhite       LicensePlateColorEnum = "white"
	LicensePlateColorEnumBlack       LicensePlateColorEnum = "black"
	LicensePlateColorEnumYellowGreen LicensePlateColorEnum = "yellow_green"
	LicensePlateColorEnumUnknow      LicensePlateColorEnum = "unknow"
	LicensePlateColorEnumPenyin      LicensePlateColorEnum = "penyin"
)

// All allowed values of LicensePlateColorEnum enum
var AllowedLicensePlateColorEnumEnumValues = []LicensePlateColorEnum{
	"blue",
	"green",
	"yellow",
	"white",
	"black",
	"yellow_green",
	"unknow",
	"penyin",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LicensePlateColorEnum) IsValid() bool {
	for _, existing := range AllowedLicensePlateColorEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
