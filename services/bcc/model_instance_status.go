package bcc

// InstanceStatus the model 'InstanceStatus'
type InstanceStatus string

// List of InstanceStatus
const (
	InstanceStatusStarting           InstanceStatus = "Starting"
	InstanceStatusRunning            InstanceStatus = "Running"
	InstanceStatusStopping           InstanceStatus = "Stopping"
	InstanceStatusStopped            InstanceStatus = "Stopped"
	InstanceStatusRecycled           InstanceStatus = "Recycled"
	InstanceStatusDeleted            InstanceStatus = "Deleted"
	InstanceStatusScaling            InstanceStatus = "Scaling"
	InstanceStatusExpired            InstanceStatus = "Expired"
	InstanceStatusError              InstanceStatus = "Error"
	InstanceStatusSnapshotprocessing InstanceStatus = "SnapshotProcessing"
	InstanceStatusImageprocessing    InstanceStatus = "ImageProcessing"
	InstanceStatusRecharging         InstanceStatus = "Recharging"
	InstanceStatusVolumeresizing     InstanceStatus = "VolumeResizing"
	InstanceStatusBillingchanging    InstanceStatus = "BillingChanging"
	InstanceStatusChangesubnet       InstanceStatus = "ChangeSubnet"
	InstanceStatusChangevpc          InstanceStatus = "ChangeVpc"
	InstanceStatusAttachingport      InstanceStatus = "AttachingPort"
	InstanceStatusDetachingport      InstanceStatus = "DetachingPort"
	InstanceStatusMoving             InstanceStatus = "Moving"
)

// All allowed values of InstanceStatus enum
var AllowedInstanceStatusEnumValues = []InstanceStatus{
	"Starting",
	"Running",
	"Stopping",
	"Stopped",
	"Recycled",
	"Deleted",
	"Scaling",
	"Expired",
	"Error",
	"SnapshotProcessing",
	"ImageProcessing",
	"Recharging",
	"VolumeResizing",
	"BillingChanging",
	"ChangeSubnet",
	"ChangeVpc",
	"AttachingPort",
	"DetachingPort",
	"Moving",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstanceStatus) IsValid() bool {
	for _, existing := range AllowedInstanceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
