package ax

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateSandboxSnapshotResponse struct {
	bce.BaseResponse
	SnapshotID *string   `json:"snapshotID,omitempty"`
	Names      []*string `json:"names,omitempty"`
}
