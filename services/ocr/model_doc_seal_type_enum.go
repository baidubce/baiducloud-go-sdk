package ocr

// DocSealTypeEnum the model 'DocSealTypeEnum'
type DocSealTypeEnum string

// List of DocSealTypeEnum
const (
	DocSealTypeEnumCircle    DocSealTypeEnum = "circle"
	DocSealTypeEnumEllipse   DocSealTypeEnum = "ellipse"
	DocSealTypeEnumRectangle DocSealTypeEnum = "rectangle"
)

// All allowed values of DocSealTypeEnum enum
var AllowedDocSealTypeEnumEnumValues = []DocSealTypeEnum{
	"circle",
	"ellipse",
	"rectangle",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocSealTypeEnum) IsValid() bool {
	for _, existing := range AllowedDocSealTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
