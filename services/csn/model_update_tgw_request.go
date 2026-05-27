package csn

type UpdateTgwRequest struct {
	CsnId       *string `json:"-"`
	TgwId       *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
