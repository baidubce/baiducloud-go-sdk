package blb

type UpdateAppBlbPolicyRequest struct {
	BlbId       *string               `json:"-"`
	ClientToken *string               `json:"-"`
	Port        *int32                `json:"port,omitempty"`
	BlbType     *string               `json:"type,omitempty"`
	PolicyList  []*AppPolicyForUpdate `json:"policyList,omitempty"`
}
