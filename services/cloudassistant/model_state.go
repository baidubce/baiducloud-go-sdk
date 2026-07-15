package cloudassistant

// State the model 'State'
type State string

// List of State
const (
	StateFailed  State = "FAILED"
	StateRunning State = "RUNNING"
	StateSuccess State = "SUCCESS"
)

// All allowed values of State enum
var AllowedStateEnumValues = []State{
	"FAILED",
	"RUNNING",
	"SUCCESS",
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
