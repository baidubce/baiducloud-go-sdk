package rds

type UpdateAutomaticStorageExpansionConfigurationRequest struct {
	InstanceId         *string `json:"-"`
	Action             *string `json:"-"`
	FreeSpaceThreshold *int32  `json:"freeSpaceThreshold,omitempty"`
	DiskMaxLimit       *int32  `json:"diskMaxLimit,omitempty"`
}
