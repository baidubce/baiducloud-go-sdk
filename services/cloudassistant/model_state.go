package cloudassistant

// State the model 'State'
type State string

// List of State
const (
	StateFailed        State = "FAILED"
	StateRunning       State = "RUNNING"
	StateSuccess       State = "SUCCESS"
	StatePartialFailed State = "PARTIAL_FAILED"
	StatePending       State = "PENDING"
)

// All allowed values of State enum
var AllowedStateEnumValues = []State{
	"FAILED",
	"RUNNING",
	"SUCCESS",
	"PARTIAL_FAILED",
	"PENDING",
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
