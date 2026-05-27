package csn

type ResizeCsnBpRequest struct {
	CsnBpId     *string `json:"-"`
	ClientToken *string `json:"-"`
	Action      *string `json:"-"`
	Bandwidth   *int32  `json:"bandwidth,omitempty"`
}
