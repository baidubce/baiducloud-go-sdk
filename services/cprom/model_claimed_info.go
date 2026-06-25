package cprom

type ClaimedInfo struct {
	IsClaimed   *bool   `json:"isClaimed,omitempty"`
	UserName    *string `json:"userName,omitempty"`
	ClaimTime   *int32  `json:"claimTime,omitempty"`
	ClaimReason *string `json:"claimReason,omitempty"`
}
