package bcc

// InstanceType the model 'InstanceType'
type InstanceType string

// List of InstanceType
const (
	InstanceTypeN1 InstanceType = "N1"
	InstanceTypeN2 InstanceType = "N2"
	InstanceTypeN3 InstanceType = "N3"
	InstanceTypeN4 InstanceType = "N4"
	InstanceTypeN5 InstanceType = "N5"
	InstanceTypeN6 InstanceType = "N6"
	InstanceTypeC1 InstanceType = "C1"
	InstanceTypeC2 InstanceType = "C2"
	InstanceTypeS1 InstanceType = "S1"
	InstanceTypeG1 InstanceType = "G1"
	InstanceTypeF1 InstanceType = "F1"
)

// All allowed values of InstanceType enum
var AllowedInstanceTypeEnumValues = []InstanceType{
	"N1",
	"N2",
	"N3",
	"N4",
	"N5",
	"N6",
	"C1",
	"C2",
	"S1",
	"G1",
	"F1",
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
