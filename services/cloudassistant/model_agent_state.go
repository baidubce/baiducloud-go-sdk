package cloudassistant

// AgentState the model 'AgentState'
type AgentState string

// List of AgentState
const (
	AgentStateOnline  AgentState = "ONLINE"
	AgentStateOffline AgentState = "OFFLINE"
)

// All allowed values of AgentState enum
var AllowedAgentStateEnumValues = []AgentState{
	"ONLINE",
	"OFFLINE",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentState) IsValid() bool {
	for _, existing := range AllowedAgentStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
