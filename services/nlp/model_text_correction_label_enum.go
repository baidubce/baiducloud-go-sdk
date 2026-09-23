package nlp

// TextCorrectionLabelEnum the model 'TextCorrectionLabelEnum'
type TextCorrectionLabelEnum string

// List of TextCorrectionLabelEnum
const (
	TextCorrectionLabelEnumValue010100 TextCorrectionLabelEnum = "010100"
	TextCorrectionLabelEnumValue010200 TextCorrectionLabelEnum = "010200"
	TextCorrectionLabelEnumValue010600 TextCorrectionLabelEnum = "010600"
	TextCorrectionLabelEnumValue020100 TextCorrectionLabelEnum = "020100"
	TextCorrectionLabelEnumValue020200 TextCorrectionLabelEnum = "020200"
	TextCorrectionLabelEnumValue020300 TextCorrectionLabelEnum = "020300"
	TextCorrectionLabelEnumValue030100 TextCorrectionLabelEnum = "030100"
	TextCorrectionLabelEnumValue030200 TextCorrectionLabelEnum = "030200"
	TextCorrectionLabelEnumValue030300 TextCorrectionLabelEnum = "030300"
	TextCorrectionLabelEnumValue030400 TextCorrectionLabelEnum = "030400"
	TextCorrectionLabelEnumValue040101 TextCorrectionLabelEnum = "040101"
	TextCorrectionLabelEnumValue040102 TextCorrectionLabelEnum = "040102"
	TextCorrectionLabelEnumValue040200 TextCorrectionLabelEnum = "040200"
	TextCorrectionLabelEnumValue040300 TextCorrectionLabelEnum = "040300"
	TextCorrectionLabelEnumValue040400 TextCorrectionLabelEnum = "040400"
	TextCorrectionLabelEnumValue060100 TextCorrectionLabelEnum = "060100"
	TextCorrectionLabelEnumValue060200 TextCorrectionLabelEnum = "060200"
	TextCorrectionLabelEnumValue060300 TextCorrectionLabelEnum = "060300"
)

// All allowed values of TextCorrectionLabelEnum enum
var AllowedTextCorrectionLabelEnumEnumValues = []TextCorrectionLabelEnum{
	"010100",
	"010200",
	"010600",
	"020100",
	"020200",
	"020300",
	"030100",
	"030200",
	"030300",
	"030400",
	"040101",
	"040102",
	"040200",
	"040300",
	"040400",
	"060100",
	"060200",
	"060300",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TextCorrectionLabelEnum) IsValid() bool {
	for _, existing := range AllowedTextCorrectionLabelEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
