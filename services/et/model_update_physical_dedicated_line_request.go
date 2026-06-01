package et

type UpdatePhysicalDedicatedLineRequest struct {
	DcphyId     *string `json:"-"`
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	UserName    *string `json:"userName,omitempty"`
	UserPhone   *string `json:"userPhone,omitempty"`
	UserEmail   *string `json:"userEmail,omitempty"`
	LinkDelay   *int32  `json:"linkDelay,omitempty"`
}
