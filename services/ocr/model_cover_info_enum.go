package ocr

// CoverInfoEnum the model 'CoverInfoEnum'
type CoverInfoEnum string

// List of CoverInfoEnum
const (
	CoverInfoEnumIncomplete CoverInfoEnum = "incomplete"
	CoverInfoEnumComplete   CoverInfoEnum = "complete"
)

// All allowed values of CoverInfoEnum enum
var AllowedCoverInfoEnumEnumValues = []CoverInfoEnum{
	"incomplete",
	"complete",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CoverInfoEnum) IsValid() bool {
	for _, existing := range AllowedCoverInfoEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
