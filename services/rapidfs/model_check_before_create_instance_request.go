package rapidfs

type CheckBeforeCreateInstanceRequest struct {
	Action             *string `json:"-"`
	Zone               *string `json:"zone,omitempty"`
	VpcId              *string `json:"vpcId,omitempty"`
	SubnetId           *string `json:"subnetId,omitempty"`
	ManagedMode        *string `json:"managedMode,omitempty"`
	MetaSpec           *string `json:"metaSpec,omitempty"`
	DataSpec           *string `json:"dataSpec,omitempty"`
	RapidfsType        *string `json:"type,omitempty"`
	CapacityTiB        *int32  `json:"capacityTiB,omitempty"`
	CceClusterId       *string `json:"cceClusterId,omitempty"`
	AihcResourcePoolId *string `json:"aihcResourcePoolId,omitempty"`
	K8sControllerId    *string `json:"k8sControllerId,omitempty"`
	K8sControllerToken *string `json:"k8sControllerToken,omitempty"`
}
