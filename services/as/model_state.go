package as

// State the model 'State'
type State string

// List of State
const (
	StateEnable  State = "ENABLE"
	StateDisable State = "DISABLE"
)

// All allowed values of State enum
var AllowedStateEnumValues = []State{
	"ENABLE",
	"DISABLE",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v State) IsValid() bool {
	for _, existing := range AllowedStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
