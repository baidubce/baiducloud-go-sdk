package ocr

// CardSideEnum the model 'CardSideEnum'
type CardSideEnum string

// List of CardSideEnum
const (
	CardSideEnumIdcardFront CardSideEnum = "idcard_front"
	CardSideEnumIdcardBack  CardSideEnum = "idcard_back"
)

// All allowed values of CardSideEnum enum
var AllowedCardSideEnumEnumValues = []CardSideEnum{
	"idcard_front",
	"idcard_back",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardSideEnum) IsValid() bool {
	for _, existing := range AllowedCardSideEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
