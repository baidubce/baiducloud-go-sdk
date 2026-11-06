package rds

type ModifyDescriptionInformationRequest struct {
	InstanceId *string `json:"-"`
	DbName     *string `json:"-"`
	Remark     *string `json:"remark,omitempty"`
}
