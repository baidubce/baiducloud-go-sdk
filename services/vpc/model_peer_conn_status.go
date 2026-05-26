package vpc

// PeerConnStatus the model 'PeerConnStatus'
type PeerConnStatus string

// List of PeerConnStatus
const (
	PeerConnStatusCreating      PeerConnStatus = "creating"
	PeerConnStatusConsulting    PeerConnStatus = "consulting"
	PeerConnStatusConsultFailed PeerConnStatus = "consult_failed"
	PeerConnStatusActive        PeerConnStatus = "active"
	PeerConnStatusDown          PeerConnStatus = "down"
	PeerConnStatusStarting      PeerConnStatus = "starting"
	PeerConnStatusStopping      PeerConnStatus = "stopping"
	PeerConnStatusDeleting      PeerConnStatus = "deleting"
	PeerConnStatusDeleted       PeerConnStatus = "deleted"
	PeerConnStatusExpired       PeerConnStatus = "expired"
	PeerConnStatusError         PeerConnStatus = "error"
	PeerConnStatusUpdating      PeerConnStatus = "updating"
)

// All allowed values of PeerConnStatus enum
var AllowedPeerConnStatusEnumValues = []PeerConnStatus{
	"creating",
	"consulting",
	"consult_failed",
	"active",
	"down",
	"starting",
	"stopping",
	"deleting",
	"deleted",
	"expired",
	"error",
	"updating",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PeerConnStatus) IsValid() bool {
	for _, existing := range AllowedPeerConnStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
