package ocr

// ForgeryDetectionResultEnum the model 'ForgeryDetectionResultEnum'
type ForgeryDetectionResultEnum string

// List of ForgeryDetectionResultEnum
const (
	ForgeryDetectionResultEnumFake ForgeryDetectionResultEnum = "fake"
	ForgeryDetectionResultEnumReal ForgeryDetectionResultEnum = "real"
)

// All allowed values of ForgeryDetectionResultEnum enum
var AllowedForgeryDetectionResultEnumEnumValues = []ForgeryDetectionResultEnum{
	"fake",
	"real",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ForgeryDetectionResultEnum) IsValid() bool {
	for _, existing := range AllowedForgeryDetectionResultEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
