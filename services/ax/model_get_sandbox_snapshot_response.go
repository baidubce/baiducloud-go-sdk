package ax

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetSandboxSnapshotResponse struct {
	bce.BaseResponse
	SandboxId *string       `json:"sandboxId,omitempty"`
	Snapshot  *SnapshotInfo `json:"snapshot,omitempty"`
}
