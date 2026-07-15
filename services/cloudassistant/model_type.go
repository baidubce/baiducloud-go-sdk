package cloudassistant

// Type the model 'Type'
type Type string

// List of Type
const (
	TypeShell      Type = "SHELL"
	TypePowershell Type = "POWERSHELL"
)

// All allowed values of Type enum
var AllowedTypeEnumValues = []Type{
	"SHELL",
	"POWERSHELL",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Type) IsValid() bool {
	for _, existing := range AllowedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
