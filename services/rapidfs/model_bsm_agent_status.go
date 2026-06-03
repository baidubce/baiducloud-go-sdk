package rapidfs

// BsmAgentStatus the model 'BsmAgentStatus'
type BsmAgentStatus string

// List of BsmAgentStatus
const (
	BsmAgentStatusOnline  BsmAgentStatus = "ONLINE"
	BsmAgentStatusOffline BsmAgentStatus = "OFFLINE"
)

// All allowed values of BsmAgentStatus enum
var AllowedBsmAgentStatusEnumValues = []BsmAgentStatus{
	"ONLINE",
	"OFFLINE",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BsmAgentStatus) IsValid() bool {
	for _, existing := range AllowedBsmAgentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
