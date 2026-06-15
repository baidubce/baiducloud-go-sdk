package bcc

type RebuildInstanceRequest struct {
	InstanceId        *string `json:"-"`
	ImageId           *string `json:"imageId,omitempty"`
	KeepImageLogin    *bool   `json:"keepImageLogin,omitempty"`
	IsPreserveData    *bool   `json:"isPreserveData,omitempty"`
	AdminPass         *string `json:"adminPass,omitempty"`
	IsOpenHostEye     *bool   `json:"isOpenHostEye,omitempty"`
	SysRootSize       *int32  `json:"sysRootSize,omitempty"`
	KeypairId         *string `json:"keypairId,omitempty"`
	DataPartitionType *string `json:"dataPartitionType,omitempty"`
	RootPartitionType *string `json:"rootPartitionType,omitempty"`
	RaidId            *string `json:"raidId,omitempty"`
	UserData          *string `json:"userData,omitempty"`
	UseLastUserData   *bool   `json:"useLastUserData,omitempty"`
	CleanLastUserData *bool   `json:"cleanLastUserData,omitempty"`
}
