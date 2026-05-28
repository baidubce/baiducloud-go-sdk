package csn

type ResizeCsnBpRequest struct {
	CsnBpId     *string `json:"-"`
	ClientToken *string `json:"-"`
	Bandwidth   *int32  `json:"bandwidth,omitempty"`
}
