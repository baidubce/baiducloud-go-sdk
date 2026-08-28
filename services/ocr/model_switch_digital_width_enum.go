package ocr

// SwitchDigitalWidthEnum the model 'SwitchDigitalWidthEnum'
type SwitchDigitalWidthEnum string

// List of SwitchDigitalWidthEnum
const (
	SwitchDigitalWidthEnumAuto SwitchDigitalWidthEnum = "auto"
	SwitchDigitalWidthEnumHalf SwitchDigitalWidthEnum = "half"
	SwitchDigitalWidthEnumFull SwitchDigitalWidthEnum = "full"
)

// All allowed values of SwitchDigitalWidthEnum enum
var AllowedSwitchDigitalWidthEnumEnumValues = []SwitchDigitalWidthEnum{
	"auto",
	"half",
	"full",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SwitchDigitalWidthEnum) IsValid() bool {
	for _, existing := range AllowedSwitchDigitalWidthEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
