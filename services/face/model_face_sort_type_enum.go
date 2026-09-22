package face

// FaceSortTypeEnum the model 'FaceSortTypeEnum'
type FaceSortTypeEnum int32

// List of FaceSortTypeEnum
const (
	FaceSortTypeEnumValue0 FaceSortTypeEnum = 0
	FaceSortTypeEnumValue1 FaceSortTypeEnum = 1
)

// All allowed values of FaceSortTypeEnum enum
var AllowedFaceSortTypeEnumEnumValues = []FaceSortTypeEnum{
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FaceSortTypeEnum) IsValid() bool {
	for _, existing := range AllowedFaceSortTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
