package ocr

// SealTypeEnum the model 'SealTypeEnum'
type SealTypeEnum string

// List of SealTypeEnum
const (
	SealTypeEnumCircle      SealTypeEnum = "circle"
	SealTypeEnumEllipse     SealTypeEnum = "ellipse"
	SealTypeEnumRectangle   SealTypeEnum = "rectangle"
	SealTypeEnumPerforation SealTypeEnum = "perforation"
)

// All allowed values of SealTypeEnum enum
var AllowedSealTypeEnumEnumValues = []SealTypeEnum{
	"circle",
	"ellipse",
	"rectangle",
	"perforation",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SealTypeEnum) IsValid() bool {
	for _, existing := range AllowedSealTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
