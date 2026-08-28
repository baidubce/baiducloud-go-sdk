package ocr

// HandwritingEngGranularityEnum the model 'HandwritingEngGranularityEnum'
type HandwritingEngGranularityEnum string

// List of HandwritingEngGranularityEnum
const (
	HandwritingEngGranularityEnumLetter HandwritingEngGranularityEnum = "letter"
	HandwritingEngGranularityEnumWord   HandwritingEngGranularityEnum = "word"
)

// All allowed values of HandwritingEngGranularityEnum enum
var AllowedHandwritingEngGranularityEnumEnumValues = []HandwritingEngGranularityEnum{
	"letter",
	"word",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HandwritingEngGranularityEnum) IsValid() bool {
	for _, existing := range AllowedHandwritingEngGranularityEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
