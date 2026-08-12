package ax

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "ax." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_SANDBOXES = "sandboxes"

	CONSTANT_RESOURCES = "resources"

	CONSTANT_TIMEOUT = "timeout"

	CONSTANT_V2 = "v2"

	CONSTANT_RESUME = "resume"

	CONSTANT_SNAPSHOTS = "snapshots"

	CONSTANT_BATCH_RELEASE = "batchRelease"

	CONSTANT_CONNECT = "connect"

	CONSTANT_FORK = "fork"

	CONSTANT_QUERY = "query"

	CONSTANT_PAUSE = "pause"
)

// Client of ax service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func NewClientWithApiKey(apiKey, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithApiKey(apiKey, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getBatchReleaseSandboxesUri() string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + CONSTANT_BATCH_RELEASE
}
func getConnectSandboxUri(SandboxID string) string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + SandboxID + bce.URI_PREFIX + CONSTANT_CONNECT
}
func getCreateSandboxUri() string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES
}
func getCreateSandboxSnapshotUri(SandboxID string) string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + SandboxID + bce.URI_PREFIX + CONSTANT_SNAPSHOTS
}
func getDeleteSandboxUri(SandboxID string) string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + SandboxID
}
func getForkSandboxUri(SandboxID string) string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + SandboxID + bce.URI_PREFIX + CONSTANT_FORK
}
func getGetSandboxUri(SandboxID string) string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + SandboxID
}
func getGetSandboxResourcesUri(SandboxID string) string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + SandboxID + bce.URI_PREFIX + CONSTANT_RESOURCES
}
func getGetSandboxSnapshotUri(SandboxID string, SnapshotID string) string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + SandboxID + bce.URI_PREFIX + CONSTANT_SNAPSHOTS + bce.URI_PREFIX + SnapshotID
}
func getListSandboxSnapshotsUri(SandboxID string) string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + SandboxID + bce.URI_PREFIX + CONSTANT_SNAPSHOTS
}
func getListSandboxesUri() string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES
}
func getListSandboxesV2Uri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_SANDBOXES
}
func getListSandboxesV2ByPathUri() string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + CONSTANT_V2
}
func getPauseSandboxUri(SandboxID string) string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + SandboxID + bce.URI_PREFIX + CONSTANT_PAUSE
}
func getQuerySandboxesUri() string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + CONSTANT_QUERY
}
func getResumeSandboxUri(SandboxID string) string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + SandboxID + bce.URI_PREFIX + CONSTANT_RESUME
}
func getSetSandboxTimeoutUri(SandboxID string) string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + SandboxID + bce.URI_PREFIX + CONSTANT_TIMEOUT
}
