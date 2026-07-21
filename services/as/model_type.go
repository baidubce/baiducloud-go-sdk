package as

// Type the model 'Type'
type Type string

// List of Type
const (
	TypeCrontab Type = "CRONTAB"
	TypePeriod  Type = "PERIOD"
	TypeAlarm   Type = "ALARM"
)

// All allowed values of Type enum
var AllowedTypeEnumValues = []Type{
	"CRONTAB",
	"PERIOD",
	"ALARM",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Type) IsValid() bool {
	for _, existing := range AllowedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
