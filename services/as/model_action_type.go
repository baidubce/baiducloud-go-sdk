package as

// ActionType the model 'ActionType'
type ActionType string

// List of ActionType
const (
	ActionTypeIncrease ActionType = "INCREASE"
	ActionTypeDecrease ActionType = "DECREASE"
	ActionTypeAdjust   ActionType = "ADJUST"
)

// All allowed values of ActionType enum
var AllowedActionTypeEnumValues = []ActionType{
	"INCREASE",
	"DECREASE",
	"ADJUST",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActionType) IsValid() bool {
	for _, existing := range AllowedActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
