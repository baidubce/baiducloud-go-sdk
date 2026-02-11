package eip

// InstanceType the model 'InstanceType'
type InstanceType string

// List of InstanceType
const (
	BCC InstanceType = "BCC"
	BBC InstanceType = "BBC"
	DCC InstanceType = "DCC"
	ENI InstanceType = "ENI"
	BLB InstanceType = "BLB"
	VPN InstanceType = "VPN"
	NAT InstanceType = "NAT"
)

// All allowed values of InstanceType enum
var AllowedInstanceTypeEnumValues = []InstanceType{
	"BCC",
	"BBC",
	"DCC",
	"ENI",
	"BLB",
	"VPN",
	"NAT",
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
