package ocr

// WordTypeEnum the model 'WordTypeEnum'
type WordTypeEnum string

// List of WordTypeEnum
const (
	WordTypeEnumHandwriting WordTypeEnum = "handwriting"
	WordTypeEnumPrint       WordTypeEnum = "print"
)

// All allowed values of WordTypeEnum enum
var AllowedWordTypeEnumEnumValues = []WordTypeEnum{
	"handwriting",
	"print",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WordTypeEnum) IsValid() bool {
	for _, existing := range AllowedWordTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
