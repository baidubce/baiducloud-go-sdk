package as

// ExpansionStrategy the model 'ExpansionStrategy'
type ExpansionStrategy string

// List of ExpansionStrategy
const (
	ExpansionStrategyPriority ExpansionStrategy = "Priority"
	ExpansionStrategyBalanced ExpansionStrategy = "Balanced"
)

// All allowed values of ExpansionStrategy enum
var AllowedExpansionStrategyEnumValues = []ExpansionStrategy{
	"Priority",
	"Balanced",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExpansionStrategy) IsValid() bool {
	for _, existing := range AllowedExpansionStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
