package cloudassistant

// Scope the model 'Scope'
type Scope string

// List of Scope
const (
	ScopeIndividual Scope = "INDIVIDUAL"
	ScopeGlobal     Scope = "GLOBAL"
)

// All allowed values of Scope enum
var AllowedScopeEnumValues = []Scope{
	"INDIVIDUAL",
	"GLOBAL",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Scope) IsValid() bool {
	for _, existing := range AllowedScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
