package ocr

// WordsAttributeEnum the model 'WordsAttributeEnum'
type WordsAttributeEnum string

// List of WordsAttributeEnum
const (
	WordsAttributeEnumHandwriting WordsAttributeEnum = "handwriting"
	WordsAttributeEnumPrint       WordsAttributeEnum = "print"
)

// All allowed values of WordsAttributeEnum enum
var AllowedWordsAttributeEnumEnumValues = []WordsAttributeEnum{
	"handwriting",
	"print",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WordsAttributeEnum) IsValid() bool {
	for _, existing := range AllowedWordsAttributeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
