package ocr

// HouseholdRegisterSideEnum the model 'HouseholdRegisterSideEnum'
type HouseholdRegisterSideEnum string

// List of HouseholdRegisterSideEnum
const (
	HouseholdRegisterSideEnumSubpage  HouseholdRegisterSideEnum = "subpage"
	HouseholdRegisterSideEnumHomepage HouseholdRegisterSideEnum = "homepage"
)

// All allowed values of HouseholdRegisterSideEnum enum
var AllowedHouseholdRegisterSideEnumEnumValues = []HouseholdRegisterSideEnum{
	"subpage",
	"homepage",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HouseholdRegisterSideEnum) IsValid() bool {
	for _, existing := range AllowedHouseholdRegisterSideEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
