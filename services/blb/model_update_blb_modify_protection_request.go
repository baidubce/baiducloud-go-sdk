package blb

type UpdateBlbModifyProtectionRequest struct {
	BlbId                        *string `json:"-"`
	ClientToken                  *string `json:"-"`
	AllowModify                  *bool   `json:"allowModify,omitempty"`
	ModificationProtectionReason *string `json:"modificationProtectionReason,omitempty"`
}
