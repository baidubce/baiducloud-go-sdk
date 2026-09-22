package face

// FaceImageTypeEnum the model 'FaceImageTypeEnum'
type FaceImageTypeEnum string

// List of FaceImageTypeEnum
const (
	FaceImageTypeEnumBase64    FaceImageTypeEnum = "BASE64"
	FaceImageTypeEnumFaceToken FaceImageTypeEnum = "FACE_TOKEN"
)

// All allowed values of FaceImageTypeEnum enum
var AllowedFaceImageTypeEnumEnumValues = []FaceImageTypeEnum{
	"BASE64",
	"FACE_TOKEN",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FaceImageTypeEnum) IsValid() bool {
	for _, existing := range AllowedFaceImageTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
