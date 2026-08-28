package ocr

// WordtypeAttriEnum the model 'WordtypeAttriEnum'
type WordtypeAttriEnum string

// List of WordtypeAttriEnum
const (
	WordtypeAttriEnumHandwriting WordtypeAttriEnum = "handwriting"
	WordtypeAttriEnumPrint       WordtypeAttriEnum = "print"
)

// All allowed values of WordtypeAttriEnum enum
var AllowedWordtypeAttriEnumEnumValues = []WordtypeAttriEnum{
	"handwriting",
	"print",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WordtypeAttriEnum) IsValid() bool {
	for _, existing := range AllowedWordtypeAttriEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
