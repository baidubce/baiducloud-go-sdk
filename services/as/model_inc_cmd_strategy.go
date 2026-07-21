package as

// IncCmdStrategy the model 'IncCmdStrategy'
type IncCmdStrategy string

// List of IncCmdStrategy
const (
	IncCmdStrategyProceed IncCmdStrategy = "Proceed"
	IncCmdStrategyPause   IncCmdStrategy = "Pause"
)

// All allowed values of IncCmdStrategy enum
var AllowedIncCmdStrategyEnumValues = []IncCmdStrategy{
	"Proceed",
	"Pause",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IncCmdStrategy) IsValid() bool {
	for _, existing := range AllowedIncCmdStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
