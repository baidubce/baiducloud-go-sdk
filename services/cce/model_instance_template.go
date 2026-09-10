package cce

type InstanceTemplate struct {
	MachineType               *string            `json:"machineType,omitempty"`
	InstanceType              *string            `json:"instanceType,omitempty"`
	InstanceName              *string            `json:"instanceName,omitempty"`
	VpcConfig                 *interface{}       `json:"vpcConfig,omitempty"`
	InstanceResource          *interface{}       `json:"instanceResource,omitempty"`
	CheckGPUDriver            *bool              `json:"checkGPUDriver,omitempty"`
	ImageID                   *string            `json:"imageID,omitempty"`
	UserData                  *interface{}       `json:"userData,omitempty"`
	InstanceOS                *interface{}       `json:"instanceOS,omitempty"`
	ScaleDownDisabled         *bool              `json:"scaleDownDisabled,omitempty"`
	IsOpenHostnameDomain      *bool              `json:"isOpenHostnameDomain,omitempty"`
	NeedEIP                   *bool              `json:"needEIP,omitempty"`
	EipOption                 *interface{}       `json:"eipOption,omitempty"`
	IamRole                   *interface{}       `json:"iamRole,omitempty"`
	DeployCustomConfig        *interface{}       `json:"deployCustomConfig,omitempty"`
	RuntimeType               *string            `json:"runtimeType,omitempty"`
	RuntimeVersion            *string            `json:"runtimeVersion,omitempty"`
	DeploySetIDs              []*string          `json:"deploySetIDs,omitempty"`
	Labels                    *map[string]string `json:"labels,omitempty"`
	Annotations               *map[string]string `json:"annotations,omitempty"`
	Tags                      []*string          `json:"tags,omitempty"`
	Taints                    []*string          `json:"taints,omitempty"`
	RelationTag               *bool              `json:"relationTag,omitempty"`
	InstancePreChargingOption *interface{}       `json:"instancePreChargingOption,omitempty"`
}
