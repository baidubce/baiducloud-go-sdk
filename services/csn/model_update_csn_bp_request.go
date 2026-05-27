package csn

type UpdateCsnBpRequest struct {
	CsnBpId     *string `json:"-"`
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
}
