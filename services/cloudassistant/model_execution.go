package cloudassistant

// Execution the model 'Execution'
type Execution string

// List of Execution
const (
	ExecutionSave       Execution = "SAVE"
	ExecutionRun        Execution = "RUN"
	ExecutionSaveAndRun Execution = "SAVE_AND_RUN"
)

// All allowed values of Execution enum
var AllowedExecutionEnumValues = []Execution{
	"SAVE",
	"RUN",
	"SAVE_AND_RUN",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Execution) IsValid() bool {
	for _, existing := range AllowedExecutionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
