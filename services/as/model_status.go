package as

// Status the model 'Status'
type Status string

// List of Status
const (
	StatusCreating      Status = "CREATING"
	StatusRunning       Status = "RUNNING"
	StatusScalingUp     Status = "SCALING_UP"
	StatusScalingDown   Status = "SCALING_DOWN"
	StatusAttachingNode Status = "ATTACHING_NODE"
	StatusDetachingNode Status = "DETACHING_NODE"
	StatusDeleting      Status = "DELETING"
	StatusBindingBlb    Status = "BINDING_BLB"
	StatusUnbindingBlb  Status = "UNBINDING_BLB"
	StatusCooldown      Status = "COOLDOWN"
	StatusPause         Status = "PAUSE"
	StatusDeleted       Status = "DELETED"
)

// All allowed values of Status enum
var AllowedStatusEnumValues = []Status{
	"CREATING",
	"RUNNING",
	"SCALING_UP",
	"SCALING_DOWN",
	"ATTACHING_NODE",
	"DETACHING_NODE",
	"DELETING",
	"BINDING_BLB",
	"UNBINDING_BLB",
	"COOLDOWN",
	"PAUSE",
	"DELETED",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Status) IsValid() bool {
	for _, existing := range AllowedStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
