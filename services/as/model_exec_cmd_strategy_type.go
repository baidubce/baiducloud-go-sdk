package as

// ExecCmdStrategyType the model 'ExecCmdStrategyType'
type ExecCmdStrategyType string

// List of ExecCmdStrategyType
const (
	ExecCmdStrategyTypeProceed ExecCmdStrategyType = "Proceed"
	ExecCmdStrategyTypePause   ExecCmdStrategyType = "Pause"
)

// All allowed values of ExecCmdStrategyType enum
var AllowedExecCmdStrategyTypeEnumValues = []ExecCmdStrategyType{
	"Proceed",
	"Pause",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExecCmdStrategyType) IsValid() bool {
	for _, existing := range AllowedExecCmdStrategyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
