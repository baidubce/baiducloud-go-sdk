package ax

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

// BatchReleaseSandboxes
//
// PARAMS:
//   - request: the arguments to BatchReleaseSandboxes
//
// RETURNS:
//   - BatchReleaseSandboxesResponse: The return type of the BatchReleaseSandboxes interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchReleaseSandboxes(request *BatchReleaseSandboxesRequest) (*BatchReleaseSandboxesResponse, error) {
	result := &BatchReleaseSandboxesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchReleaseSandboxesUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ConnectSandbox
//
// PARAMS:
//   - request: the arguments to ConnectSandbox
//
// RETURNS:
//   - ConnectSandboxResponse: The return type of the ConnectSandbox interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ConnectSandbox(request *ConnectSandboxRequest) (*ConnectSandboxResponse, error) {
	result := &ConnectSandboxResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getConnectSandboxUri(util.StringValue(request.SandboxID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateSandbox
//
// PARAMS:
//   - request: the arguments to CreateSandbox
//
// RETURNS:
//   - CreateSandboxResponse: The return type of the CreateSandbox interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSandbox(request *CreateSandboxRequest) (*CreateSandboxResponse, error) {
	result := &CreateSandboxResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateSandboxUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateSandboxSnapshot
//
// PARAMS:
//   - request: the arguments to CreateSandboxSnapshot
//
// RETURNS:
//   - CreateSandboxSnapshotResponse: The return type of the CreateSandboxSnapshot interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSandboxSnapshot(request *CreateSandboxSnapshotRequest) (*CreateSandboxSnapshotResponse, error) {
	result := &CreateSandboxSnapshotResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateSandboxSnapshotUri(util.StringValue(request.SandboxID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteSandbox
//
// PARAMS:
//   - request: the arguments to DeleteSandbox
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSandbox(request *DeleteSandboxRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSandboxUri(util.StringValue(request.SandboxID))).
		Do()
}

// ForkSandbox
//
// PARAMS:
//   - request: the arguments to ForkSandbox
//
// RETURNS:
//   - ForkSandboxResponse: The return type of the ForkSandbox interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ForkSandbox(request *ForkSandboxRequest) (*ForkSandboxResponse, error) {
	result := &ForkSandboxResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getForkSandboxUri(util.StringValue(request.SandboxID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSandbox
//
// PARAMS:
//   - request: the arguments to GetSandbox
//
// RETURNS:
//   - GetSandboxResponse: The return type of the GetSandbox interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetSandbox(request *GetSandboxRequest) (*GetSandboxResponse, error) {
	result := &GetSandboxResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetSandboxUri(util.StringValue(request.SandboxID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSandboxResources
//
// PARAMS:
//   - request: the arguments to GetSandboxResources
//
// RETURNS:
//   - GetSandboxResourcesResponse: The return type of the GetSandboxResources interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetSandboxResources(request *GetSandboxResourcesRequest) (*GetSandboxResourcesResponse, error) {
	result := &GetSandboxResourcesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetSandboxResourcesUri(util.StringValue(request.SandboxID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSandboxSnapshot
//
// PARAMS:
//   - request: the arguments to GetSandboxSnapshot
//
// RETURNS:
//   - GetSandboxSnapshotResponse: The return type of the GetSandboxSnapshot interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetSandboxSnapshot(request *GetSandboxSnapshotRequest) (*GetSandboxSnapshotResponse, error) {
	result := &GetSandboxSnapshotResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetSandboxSnapshotUri(util.StringValue(request.SandboxID), util.StringValue(request.SnapshotID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListSandboxSnapshots
//
// PARAMS:
//   - request: the arguments to ListSandboxSnapshots
//
// RETURNS:
//   - ListSandboxSnapshotsResponse: The return type of the ListSandboxSnapshots interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListSandboxSnapshots(request *ListSandboxSnapshotsRequest) (*ListSandboxSnapshotsResponse, error) {
	result := &ListSandboxSnapshotsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListSandboxSnapshotsUri(util.StringValue(request.SandboxID))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListSandboxes
//
// PARAMS:
//   - request: the arguments to ListSandboxes
//
// RETURNS:
//   - ListSandboxesResponse: The return type of the ListSandboxes interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListSandboxes(request *ListSandboxesRequest) (*ListSandboxesResponse, error) {
	result := &ListSandboxesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListSandboxesUri()).
		WithQueryParamFilter("metadata", util.StringValue(request.Metadata)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListSandboxesV2
//
// PARAMS:
//   - request: the arguments to ListSandboxesV2
//
// RETURNS:
//   - ListSandboxesV2Response: The return type of the ListSandboxesV2 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListSandboxesV2(request *ListSandboxesV2Request) (*ListSandboxesV2Response, error) {
	result := &ListSandboxesV2Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListSandboxesV2Uri()).
		WithQueryParamFilter("limit", util.Int32Value(request.Limit)).
		WithQueryParamFilter("nextToken", util.StringValue(request.NextToken)).
		WithQueryParamFilter("metadata", util.StringValue(request.Metadata)).
		WithQueryParamFilter("state", util.StringValue(request.State)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListSandboxesV2ByPath
//
// PARAMS:
//   - request: the arguments to ListSandboxesV2ByPath
//
// RETURNS:
//   - ListSandboxesV2ByPathResponse: The return type of the ListSandboxesV2ByPath interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListSandboxesV2ByPath(request *ListSandboxesV2ByPathRequest) (*ListSandboxesV2ByPathResponse, error) {
	result := &ListSandboxesV2ByPathResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListSandboxesV2ByPathUri()).
		WithQueryParamFilter("limit", util.Int32Value(request.Limit)).
		WithQueryParamFilter("nextToken", util.StringValue(request.NextToken)).
		WithQueryParamFilter("metadata", util.StringValue(request.Metadata)).
		WithQueryParamFilter("state", util.StringValue(request.State)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PauseSandbox
//
// PARAMS:
//   - request: the arguments to PauseSandbox
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) PauseSandbox(request *PauseSandboxRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPauseSandboxUri(util.StringValue(request.SandboxID))).
		WithQueryParamFilter("hibernateMode", util.StringValue(request.HibernateMode)).
		Do()
}

// QuerySandboxes
//
// PARAMS:
//   - request: the arguments to QuerySandboxes
//
// RETURNS:
//   - QuerySandboxesResponse: The return type of the QuerySandboxes interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QuerySandboxes(request *QuerySandboxesRequest) (*QuerySandboxesResponse, error) {
	result := &QuerySandboxesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getQuerySandboxesUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResumeSandbox
//
// PARAMS:
//   - request: the arguments to ResumeSandbox
//
// RETURNS:
//   - ResumeSandboxResponse: The return type of the ResumeSandbox interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ResumeSandbox(request *ResumeSandboxRequest) (*ResumeSandboxResponse, error) {
	result := &ResumeSandboxResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getResumeSandboxUri(util.StringValue(request.SandboxID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SetSandboxTimeout
//
// PARAMS:
//   - request: the arguments to SetSandboxTimeout
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) SetSandboxTimeout(request *SetSandboxTimeoutRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSetSandboxTimeoutUri(util.StringValue(request.SandboxID))).
		WithBody(request).
		Do()
}
