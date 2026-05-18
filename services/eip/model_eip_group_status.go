package eip

// EipGroupStatus the model 'EipGroupStatus'
type EipGroupStatus string

// List of EipGroupStatus
const (
	EipGroupStatusAvailable EipGroupStatus = "available"
	EipGroupStatusPaused    EipGroupStatus = "paused"
	EipGroupStatusExpired   EipGroupStatus = "expired"
	EipGroupStatusDeleting  EipGroupStatus = "deleting"
)

// All allowed values of EipGroupStatus enum
var AllowedEipGroupStatusEnumValues = []EipGroupStatus{
	"available",
	"paused",
	"expired",
	"deleting",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EipGroupStatus) IsValid() bool {
	for _, existing := range AllowedEipGroupStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
