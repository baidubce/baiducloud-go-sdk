

package eip




type TransferEipVo struct {
	TransferId *string `json:"transferId,omitempty"`
	TransferType *string `json:"transferType,omitempty"`
	UserId *string `json:"userId,omitempty"`
	ToUserId *string `json:"toUserId,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	InstanceName *string `json:"instanceName,omitempty"`
	InstanceIP *string `json:"instanceIP,omitempty"`
	InstanceType *string `json:"instanceType,omitempty"`
	InstanceBandwidth *string `json:"instanceBandwidth,omitempty"`
	Status *string `json:"status,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`
}

