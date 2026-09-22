package face

// FaceControlLevelEnum the model 'FaceControlLevelEnum'
type FaceControlLevelEnum string

// List of FaceControlLevelEnum
const (
	FaceControlLevelEnumNone   FaceControlLevelEnum = "NONE"
	FaceControlLevelEnumLow    FaceControlLevelEnum = "LOW"
	FaceControlLevelEnumNormal FaceControlLevelEnum = "NORMAL"
	FaceControlLevelEnumHigh   FaceControlLevelEnum = "HIGH"
)

// All allowed values of FaceControlLevelEnum enum
var AllowedFaceControlLevelEnumEnumValues = []FaceControlLevelEnum{
	"NONE",
	"LOW",
	"NORMAL",
	"HIGH",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FaceControlLevelEnum) IsValid() bool {
	for _, existing := range AllowedFaceControlLevelEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
