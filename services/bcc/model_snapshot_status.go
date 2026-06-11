package bcc

// SnapshotStatus the model 'SnapshotStatus'
type SnapshotStatus string

// List of SnapshotStatus
const (
	SnapshotStatusCreating      SnapshotStatus = "Creating"
	SnapshotStatusCreatedfailed SnapshotStatus = "CreatedFailed"
	SnapshotStatusAvailable     SnapshotStatus = "Available"
	SnapshotStatusNotavailable  SnapshotStatus = "NotAvailable"
)

// All allowed values of SnapshotStatus enum
var AllowedSnapshotStatusEnumValues = []SnapshotStatus{
	"Creating",
	"CreatedFailed",
	"Available",
	"NotAvailable",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SnapshotStatus) IsValid() bool {
	for _, existing := range AllowedSnapshotStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
