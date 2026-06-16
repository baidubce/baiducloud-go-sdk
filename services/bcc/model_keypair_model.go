package bcc

type KeypairModel struct {
	KeypairId         *string              `json:"keypairId,omitempty"`
	Name              *string              `json:"name,omitempty"`
	Description       *string              `json:"description,omitempty"`
	InstanceCount     *int32               `json:"instanceCount,omitempty"`
	CreatedTime       *string              `json:"createdTime,omitempty"`
	PublicKey         *string              `json:"publicKey,omitempty"`
	FingerPrint       *string              `json:"fingerPrint,omitempty"`
	PrivateKey        *string              `json:"privateKey,omitempty"`
	RegionId          *string              `json:"regionId,omitempty"`
	PaasInstanceCount []*PassInstanceModel `json:"paasInstanceCount,omitempty"`
}
