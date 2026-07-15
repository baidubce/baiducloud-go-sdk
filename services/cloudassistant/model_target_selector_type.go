package cloudassistant

// TargetSelectorType the model 'TargetSelectorType'
type TargetSelectorType string

// List of TargetSelectorType
const (
	TargetSelectorTypeInstancesList   TargetSelectorType = "INSTANCES_LIST"
	TargetSelectorTypeAllInstances    TargetSelectorType = "ALL_INSTANCES"
	TargetSelectorTypeTagInstances    TargetSelectorType = "TAG_INSTANCES"
	TargetSelectorTypeInstancesImport TargetSelectorType = "INSTANCES_IMPORT"
)

// All allowed values of TargetSelectorType enum
var AllowedTargetSelectorTypeEnumValues = []TargetSelectorType{
	"INSTANCES_LIST",
	"ALL_INSTANCES",
	"TAG_INSTANCES",
	"INSTANCES_IMPORT",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TargetSelectorType) IsValid() bool {
	for _, existing := range AllowedTargetSelectorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
