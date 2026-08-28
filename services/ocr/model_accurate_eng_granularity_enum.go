package ocr

// AccurateEngGranularityEnum the model 'AccurateEngGranularityEnum'
type AccurateEngGranularityEnum string

// List of AccurateEngGranularityEnum
const (
	AccurateEngGranularityEnumLetter AccurateEngGranularityEnum = "letter"
	AccurateEngGranularityEnumWord   AccurateEngGranularityEnum = "word"
)

// All allowed values of AccurateEngGranularityEnum enum
var AllowedAccurateEngGranularityEnumEnumValues = []AccurateEngGranularityEnum{
	"letter",
	"word",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccurateEngGranularityEnum) IsValid() bool {
	for _, existing := range AllowedAccurateEngGranularityEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
