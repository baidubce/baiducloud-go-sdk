package blb

type DeleteAppBlbPolicyRequest struct {
	BlbId        *string   `json:"-"`
	ClientToken  *string   `json:"-"`
	Port         *int32    `json:"port,omitempty"`
	PolicyIdList []*string `json:"policyIdList,omitempty"`
	BlbType      *string   `json:"type,omitempty"`
}
