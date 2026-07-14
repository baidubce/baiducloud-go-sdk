package as

type BlbInfo struct {
	BlbId   *string   `json:"blbId,omitempty"`
	BlbName *string   `json:"blbName,omitempty"`
	BlbType *string   `json:"blbType,omitempty"`
	SgIds   []*string `json:"sgIds,omitempty"`
}
