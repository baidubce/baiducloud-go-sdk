package face

// FaceRegisterActionEnum the model 'FaceRegisterActionEnum'
type FaceRegisterActionEnum string

// List of FaceRegisterActionEnum
const (
	FaceRegisterActionEnumAppend  FaceRegisterActionEnum = "APPEND"
	FaceRegisterActionEnumReplace FaceRegisterActionEnum = "REPLACE"
)

// All allowed values of FaceRegisterActionEnum enum
var AllowedFaceRegisterActionEnumEnumValues = []FaceRegisterActionEnum{
	"APPEND",
	"REPLACE",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FaceRegisterActionEnum) IsValid() bool {
	for _, existing := range AllowedFaceRegisterActionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
