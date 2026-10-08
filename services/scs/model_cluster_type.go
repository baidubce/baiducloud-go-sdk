package scs

// ClusterType the model 'ClusterType'
type ClusterType string

// List of ClusterType
const (
	ClusterTypeCluster     ClusterType = "cluster"
	ClusterTypeMasterSlave ClusterType = "master_slave"
)

// All allowed values of ClusterType enum
var AllowedClusterTypeEnumValues = []ClusterType{
	"cluster",
	"master_slave",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterType) IsValid() bool {
	for _, existing := range AllowedClusterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
