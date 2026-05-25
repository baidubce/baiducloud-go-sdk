package vpc

// NatStatus the model 'NatStatus'
type NatStatus string

// List of NatStatus
const (
	NatStatusActive       NatStatus = "active"
	NatStatusUpdating     NatStatus = "updating"
	NatStatusUnconfigured NatStatus = "unconfigured"
	NatStatusDown         NatStatus = "down"
	NatStatusBuilding     NatStatus = "building"
	NatStatusError        NatStatus = "error"
	NatStatusDeleting     NatStatus = "deleting"
	NatStatusDeleted      NatStatus = "deleted"
	NatStatusStarting     NatStatus = "starting"
	NatStatusConfiguring  NatStatus = "configuring"
	NatStatusRebooting    NatStatus = "rebooting"
	NatStatusStopping     NatStatus = "stopping"
)

// All allowed values of NatStatus enum
var AllowedNatStatusEnumValues = []NatStatus{
	"active",
	"updating",
	"unconfigured",
	"down",
	"building",
	"error",
	"deleting",
	"deleted",
	"starting",
	"configuring",
	"rebooting",
	"stopping",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NatStatus) IsValid() bool {
	for _, existing := range AllowedNatStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
