package rds

type ModifyDatabasePortRequest struct {
	InstanceId *string `json:"-"`
	EntryPort  *int32  `json:"entryPort,omitempty"`
}
