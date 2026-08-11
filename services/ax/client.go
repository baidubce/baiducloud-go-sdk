package ax

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "ax." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_SANDBOXES = "sandboxes"

	CONSTANT_QUERY = "query"
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

func getQuerySandboxesUri() string {
	return bce.URI_PREFIX + CONSTANT_SANDBOXES + bce.URI_PREFIX + CONSTANT_QUERY
}
