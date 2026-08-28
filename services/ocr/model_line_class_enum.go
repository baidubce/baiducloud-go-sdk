package ocr

// LineClassEnum the model 'LineClassEnum'
type LineClassEnum string

// List of LineClassEnum
const (
	LineClassEnumKey        LineClassEnum = "key"
	LineClassEnumValue      LineClassEnum = "value"
	LineClassEnumTableValue LineClassEnum = "table_value"
	LineClassEnumOther      LineClassEnum = "other"
)

// All allowed values of LineClassEnum enum
var AllowedLineClassEnumEnumValues = []LineClassEnum{
	"key",
	"value",
	"table_value",
	"other",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LineClassEnum) IsValid() bool {
	for _, existing := range AllowedLineClassEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
