package cloudassistant

// Os the model 'Os'
type Os string

// List of Os
const (
	OsLinux   Os = "LINUX"
	OsWindows Os = "WINDOWS"
)

// All allowed values of Os enum
var AllowedOsEnumValues = []Os{
	"LINUX",
	"WINDOWS",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Os) IsValid() bool {
	for _, existing := range AllowedOsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
