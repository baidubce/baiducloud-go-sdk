package face

// FaceMergeDegreeEnum the model 'FaceMergeDegreeEnum'
type FaceMergeDegreeEnum string

// List of FaceMergeDegreeEnum
const (
	FaceMergeDegreeEnumLow      FaceMergeDegreeEnum = "LOW"
	FaceMergeDegreeEnumNormal   FaceMergeDegreeEnum = "NORMAL"
	FaceMergeDegreeEnumHigh     FaceMergeDegreeEnum = "HIGH"
	FaceMergeDegreeEnumComplete FaceMergeDegreeEnum = "COMPLETE"
)

// All allowed values of FaceMergeDegreeEnum enum
var AllowedFaceMergeDegreeEnumEnumValues = []FaceMergeDegreeEnum{
	"LOW",
	"NORMAL",
	"HIGH",
	"COMPLETE",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FaceMergeDegreeEnum) IsValid() bool {
	for _, existing := range AllowedFaceMergeDegreeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
