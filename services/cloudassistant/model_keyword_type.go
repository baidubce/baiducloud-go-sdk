package cloudassistant

// KeywordType the model 'KeywordType'
type KeywordType string

// List of KeywordType
const (
	KeywordTypeInstanceid KeywordType = "instanceId"
	KeywordTypeInternalip KeywordType = "internalIp"
)

// All allowed values of KeywordType enum
var AllowedKeywordTypeEnumValues = []KeywordType{
	"instanceId",
	"internalIp",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordType) IsValid() bool {
	for _, existing := range AllowedKeywordTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
