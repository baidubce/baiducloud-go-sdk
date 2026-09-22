package face

// FaceUpdateActionEnum the model 'FaceUpdateActionEnum'
type FaceUpdateActionEnum string

// List of FaceUpdateActionEnum
const (
	FaceUpdateActionEnumUpdate  FaceUpdateActionEnum = "UPDATE"
	FaceUpdateActionEnumReplace FaceUpdateActionEnum = "REPLACE"
)

// All allowed values of FaceUpdateActionEnum enum
var AllowedFaceUpdateActionEnumEnumValues = []FaceUpdateActionEnum{
	"UPDATE",
	"REPLACE",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FaceUpdateActionEnum) IsValid() bool {
	for _, existing := range AllowedFaceUpdateActionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
