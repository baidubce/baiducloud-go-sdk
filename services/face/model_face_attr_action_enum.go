package face

// FaceAttrActionEnum the model 'FaceAttrActionEnum'
type FaceAttrActionEnum string

// List of FaceAttrActionEnum
const (
	FaceAttrActionEnumToKid    FaceAttrActionEnum = "TO_KID"
	FaceAttrActionEnumToOld    FaceAttrActionEnum = "TO_OLD"
	FaceAttrActionEnumToFemale FaceAttrActionEnum = "TO_FEMALE"
	FaceAttrActionEnumToMale   FaceAttrActionEnum = "TO_MALE"
	FaceAttrActionEnumV2Age    FaceAttrActionEnum = "V2_AGE"
	FaceAttrActionEnumV2Gender FaceAttrActionEnum = "V2_GENDER"
)

// All allowed values of FaceAttrActionEnum enum
var AllowedFaceAttrActionEnumEnumValues = []FaceAttrActionEnum{
	"TO_KID",
	"TO_OLD",
	"TO_FEMALE",
	"TO_MALE",
	"V2_AGE",
	"V2_GENDER",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FaceAttrActionEnum) IsValid() bool {
	for _, existing := range AllowedFaceAttrActionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
