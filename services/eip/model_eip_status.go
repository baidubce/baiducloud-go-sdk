package eip

// EipStatus the model 'EipStatus'
type EipStatus string

// List of EipStatus
const (
	CREATING    EipStatus = "creating"
	AVAILABLE   EipStatus = "available"
	BINDED      EipStatus = "binded"
	BINDING     EipStatus = "binding"
	UNBINDING   EipStatus = "unbinding"
	UPDATING    EipStatus = "updating"
	PAUSED      EipStatus = "paused"
	UNAVAILABLE EipStatus = "unavailable"
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
