package bcc

// VolumeStatus the model 'VolumeStatus'
type VolumeStatus string

// List of VolumeStatus
const (
	VolumeStatusCreating           VolumeStatus = "Creating"
	VolumeStatusAvailable          VolumeStatus = "Available"
	VolumeStatusAttaching          VolumeStatus = "Attaching"
	VolumeStatusNotavailable       VolumeStatus = "NotAvailable"
	VolumeStatusInuse              VolumeStatus = "InUse"
	VolumeStatusDetaching          VolumeStatus = "Detaching"
	VolumeStatusDeleting           VolumeStatus = "Deleting"
	VolumeStatusDeleted            VolumeStatus = "Deleted"
	VolumeStatusScaling            VolumeStatus = "Scaling"
	VolumeStatusExpired            VolumeStatus = "Expired"
	VolumeStatusError              VolumeStatus = "Error"
	VolumeStatusSnapshotprocessing VolumeStatus = "SnapshotProcessing"
	VolumeStatusImageprocessing    VolumeStatus = "ImageProcessing"
	VolumeStatusRecharging         VolumeStatus = "Recharging"
)

// All allowed values of VolumeStatus enum
var AllowedVolumeStatusEnumValues = []VolumeStatus{
	"Creating",
	"Available",
	"Attaching",
	"NotAvailable",
	"InUse",
	"Detaching",
	"Deleting",
	"Deleted",
	"Scaling",
	"Expired",
	"Error",
	"SnapshotProcessing",
	"ImageProcessing",
	"Recharging",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VolumeStatus) IsValid() bool {
	for _, existing := range AllowedVolumeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
