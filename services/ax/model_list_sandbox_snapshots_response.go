package ax

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListSandboxSnapshotsResponse struct {
	bce.BaseResponse
	SandboxId *string         `json:"sandboxId,omitempty"`
	Snapshots []*SnapshotInfo `json:"snapshots,omitempty"`
}
