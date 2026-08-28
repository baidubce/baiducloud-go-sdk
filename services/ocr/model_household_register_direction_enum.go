package ocr

// HouseholdRegisterDirectionEnum the model 'HouseholdRegisterDirectionEnum'
type HouseholdRegisterDirectionEnum int32

// List of HouseholdRegisterDirectionEnum
const (
	HouseholdRegisterDirectionEnumValueMinus1 HouseholdRegisterDirectionEnum = -1
	HouseholdRegisterDirectionEnumValue0      HouseholdRegisterDirectionEnum = 0
	HouseholdRegisterDirectionEnumValue1      HouseholdRegisterDirectionEnum = 1
	HouseholdRegisterDirectionEnumValue2      HouseholdRegisterDirectionEnum = 2
	HouseholdRegisterDirectionEnumValue3      HouseholdRegisterDirectionEnum = 3
)

// All allowed values of HouseholdRegisterDirectionEnum enum
var AllowedHouseholdRegisterDirectionEnumEnumValues = []HouseholdRegisterDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HouseholdRegisterDirectionEnum) IsValid() bool {
	for _, existing := range AllowedHouseholdRegisterDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
