package eip

// EipInstanceType the model 'EipInstanceType'
type EipInstanceType string

// List of EipInstanceType
const (
	NORMAL EipInstanceType = "normal"
	SHARED EipInstanceType = "shared"
)

// All allowed values of EipInstanceType enum
var AllowedEipInstanceTypeEnumValues = []EipInstanceType{
	"normal",
	"shared",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EipInstanceType) IsValid() bool {
	for _, existing := range AllowedEipInstanceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
