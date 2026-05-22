package blb

// BlbStatus the model 'BlbStatus'
type BlbStatus string

// List of BlbStatus
const (
	BlbStatusCreating    BlbStatus = "creating"
	BlbStatusAvailable   BlbStatus = "available"
	BlbStatusUpdating    BlbStatus = "updating"
	BlbStatusPaused      BlbStatus = "paused"
	BlbStatusUnavailable BlbStatus = "unavailable"
)

// All allowed values of BlbStatus enum
var AllowedBlbStatusEnumValues = []BlbStatus{
	"creating",
	"available",
	"updating",
	"paused",
	"unavailable",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BlbStatus) IsValid() bool {
	for _, existing := range AllowedBlbStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
