package bcc

// VolumeType the model 'VolumeType'
type VolumeType string

// List of VolumeType
const (
	VolumeTypeSystem    VolumeType = "System"
	VolumeTypeEphemeral VolumeType = "Ephemeral"
	VolumeTypeCds       VolumeType = "Cds"
)

// All allowed values of VolumeType enum
var AllowedVolumeTypeEnumValues = []VolumeType{
	"System",
	"Ephemeral",
	"Cds",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VolumeType) IsValid() bool {
	for _, existing := range AllowedVolumeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
