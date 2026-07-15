package cloudassistant

// Id the model 'Id'
type Id string

// List of Id
const (
	IdCommand    Id = "COMMAND"
	IdFileUpload Id = "FILE_UPLOAD"
)

// All allowed values of Id enum
var AllowedIdEnumValues = []Id{
	"COMMAND",
	"FILE_UPLOAD",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Id) IsValid() bool {
	for _, existing := range AllowedIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
