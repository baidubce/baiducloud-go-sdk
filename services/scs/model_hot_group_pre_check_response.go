package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type HotGroupPreCheckResponse struct {
	bce.BaseResponse
	ConnectionResults []*ConnectionResults `json:"connectionResults,omitempty"`
	LeaderResult      *LeaderResult        `json:"leaderResult,omitempty"`
	FollowerResult    []*FollowerResult    `json:"followerResult,omitempty"`
}
