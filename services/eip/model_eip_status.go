package eip

// EipStatus the model 'EipStatus'
type EipStatus string

// List of EipStatus
const (
	EipStatusCreating    EipStatus = "creating"
	EipStatusAvailable   EipStatus = "available"
	EipStatusBinded      EipStatus = "binded"
	EipStatusBinding     EipStatus = "binding"
	EipStatusUnbinding   EipStatus = "unbinding"
	EipStatusUpdating    EipStatus = "updating"
	EipStatusPaused      EipStatus = "paused"
	EipStatusUnavailable EipStatus = "unavailable"
)

// All allowed values of EipStatus enum
var AllowedEipStatusEnumValues = []EipStatus{
	"creating",
	"available",
	"binded",
	"binding",
	"unbinding",
	"updating",
	"paused",
	"unavailable",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EipStatus) IsValid() bool {
	for _, existing := range AllowedEipStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
