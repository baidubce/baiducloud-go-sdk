package cloudassistant

// CommandType the model 'CommandType'
type CommandType string

// List of CommandType
const (
	CommandTypeShell      CommandType = "SHELL"
	CommandTypePowershell CommandType = "POWERSHELL"
)

// All allowed values of CommandType enum
var AllowedCommandTypeEnumValues = []CommandType{
	"SHELL",
	"POWERSHELL",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommandType) IsValid() bool {
	for _, existing := range AllowedCommandTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
