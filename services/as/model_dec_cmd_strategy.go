package as

// DecCmdStrategy the model 'DecCmdStrategy'
type DecCmdStrategy string

// List of DecCmdStrategy
const (
	DecCmdStrategyProceed DecCmdStrategy = "Proceed"
	DecCmdStrategyPause   DecCmdStrategy = "Pause"
)

// All allowed values of DecCmdStrategy enum
var AllowedDecCmdStrategyEnumValues = []DecCmdStrategy{
	"Proceed",
	"Pause",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DecCmdStrategy) IsValid() bool {
	for _, existing := range AllowedDecCmdStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
