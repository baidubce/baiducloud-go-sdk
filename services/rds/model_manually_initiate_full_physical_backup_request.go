package rds

type ManuallyInitiateFullPhysicalBackupRequest struct {
	InstanceId     *string `json:"-"`
	EffectiveTime  *string `json:"effectiveTime,omitempty"`
	DataBackupType *string `json:"dataBackupType,omitempty"`
}
