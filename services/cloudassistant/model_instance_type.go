package cloudassistant

// InstanceType the model 'InstanceType'
type InstanceType string

// List of InstanceType
const (
	InstanceTypeBcc InstanceType = "BCC"
	InstanceTypeBbc InstanceType = "BBC"
)

// All allowed values of InstanceType enum
var AllowedInstanceTypeEnumValues = []InstanceType{
	"BCC",
	"BBC",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstanceType) IsValid() bool {
	for _, existing := range AllowedInstanceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
