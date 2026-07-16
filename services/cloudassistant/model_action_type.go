package cloudassistant

// ActionType the model 'ActionType'
type ActionType string

// List of ActionType
const (
	ActionTypeCommand    ActionType = "COMMAND"
	ActionTypeFileUpload ActionType = "FILE_UPLOAD"
)

// All allowed values of ActionType enum
var AllowedActionTypeEnumValues = []ActionType{
	"COMMAND",
	"FILE_UPLOAD",
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
