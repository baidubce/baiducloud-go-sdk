package ocr

// IdCardSideEnum the model 'IdCardSideEnum'
type IdCardSideEnum string

// List of IdCardSideEnum
const (
	IdCardSideEnumFront IdCardSideEnum = "front"
	IdCardSideEnumBack  IdCardSideEnum = "back"
)

// All allowed values of IdCardSideEnum enum
var AllowedIdCardSideEnumEnumValues = []IdCardSideEnum{
	"front",
	"back",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdCardSideEnum) IsValid() bool {
	for _, existing := range AllowedIdCardSideEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
